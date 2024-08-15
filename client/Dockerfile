# Dockerfile - Frontend Ottimizzato

# Fase 1: Costruzione dell'applicazione
FROM node:20 AS build

# Imposta la directory di lavoro all'interno del container
WORKDIR /app

# Copia solo i file necessari per installare le dipendenze
COPY package*.json ./

# Installa le dipendenze
RUN npm ci

# Copia il resto dei file dell'applicazione
COPY . .

# Costruisci l'applicazione per la produzione
RUN npm run build

# Fase 2: Creazione dell'immagine finale
FROM node:20-slim

# Imposta la directory di lavoro all'interno del container
WORKDIR /app

# Copia solo i file necessari dall'immagine di build
COPY --from=build /app/ ./

CMD ["npm", "run", "dev", "--", "--host"]