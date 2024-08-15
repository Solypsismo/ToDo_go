package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	// "encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	Title string `json:"title"`
}

func main() {
	r := gin.Default()

	// Middleware per gestire CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Permette tutte le origini
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("errore di connessione : ", err)
	}

	fmt.Println("connessione riuscita")

	database := client.Database("testdb")
	taskCollection := database.Collection("tasks")

	r.POST("/api/task", func(c *gin.Context) {
		var tasks []Task

		if err := c.BindJSON(&tasks); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		var docs []interface{}

		for _, task := range tasks {
			docs = append(docs, task)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		result, err := taskCollection.InsertMany(ctx, docs)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		c.JSON(200, gin.H{
			"success":      true,
			"inserted_ids": result.InsertedIDs,
		})
	})

	r.GET("/api/task", func(c *gin.Context) {
		var tasks []Task

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cursor, err := taskCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &tasks); err != nil {
			c.JSON(500, gin.H{
				"success": false,
				"message": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, tasks)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	r.Run()
}
