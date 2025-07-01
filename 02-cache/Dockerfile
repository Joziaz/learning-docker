FROM golang:alpine AS builder

WORKDIR /src

# Cada commando crea un nuevo layer por lo que es mejor agruparlos
#COPY . .
COPY main.go .
COPY go.mod .

RUN go build -o ./app.exe 

