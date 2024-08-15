# Dockerfile - Backend

# Fase 1: Costruzione dell'applicazione
FROM golang:1.23-alpine AS builder

# Imposta la directory di lavoro all'interno del container
WORKDIR /app

# Copia il file go.mod e go.sum per gestire le dipendenze
COPY go.mod go.sum ./

# Scarica le dipendenze
RUN go mod download

# Copia tutto il codice sorgente nella directory di lavoro
COPY . .

# Compila l'applicazione
RUN go build -o main .

# Fase 2: Creazione dell'immagine finale
FROM alpine:latest

# Imposta la directory di lavoro
WORKDIR /root/

# Copia l'applicazione compilata dalla fase 1
COPY --from=builder /app/main .

# Esponi la porta su cui l'applicazione sarà in ascolto
EXPOSE 8080

# Comando per avviare l'applicazione
CMD ["./main"]
