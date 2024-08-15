package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Status string `json:"status,omitempty"`
	Title  string `json:"title"`
}

var tasksCollection *mongo.Collection
var client *mongo.Client

func main() {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")

	client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("errore di connessione : ", err)
	}

	fmt.Println("connessione riuscita")

	database := client.Database("testdb")
	tasksCollection = database.Collection("tasks")

	http.HandleFunc("/api/tasks", handleTask)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		getTasks(w, r)
	} else if r.Method == http.MethodPost {
		createTask(w, r)
	} else if r.Method == http.MethodDelete {
		deleteTask(w, r)
	} else if r.Method == http.MethodPut {
		doneTask(w, r)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task

	cursor, err := tasksCollection.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task Task
		if err := cursor.Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the ObjectID to a hex string
		// task.ID = task.ID.(primitive.ObjectID).Hex()

		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// settiamo a true lo status
	// TODO : andra fatto lato frontend questo passaggio
	task.Status = "in progress"

	result, err := tasksCollection.InsertOne(context.Background(), task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	task.ID = result.InsertedID.(primitive.ObjectID).Hex()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID string `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = tasksCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"deleted_id": objectID.Hex()})
}

func doneTask(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID string `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	update := bson.M{
		"$set": bson.M{
			"status": "done",
		},
	}

	_, err = tasksCollection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"updated_id": payload.ID})

}
