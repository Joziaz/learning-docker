FROM golang:alpine AS builder

WORKDIR /go/src

COPY . .

RUN go build -o ./app.exe 

# Es mejor crear una segunda imagen para reducir el tamaño final
# de la imagen y evitar que se incluya el compilador de Go
FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /go/src/app.exe ./

CMD [ "./app.exe" ]