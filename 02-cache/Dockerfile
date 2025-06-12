FROM golang:alpine AS builder

WORKDIR /src

# Cada commando crea un nuevo layer por lo que es mejor agruparlos
COPY . .

RUN go build -o ./app.exe 

