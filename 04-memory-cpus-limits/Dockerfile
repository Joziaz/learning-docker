FROM golang:alpine AS builder

WORKDIR /go/src

COPY . .

RUN go build -o ./app.exe 


FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /go/src/app.exe ./

CMD [ "./app.exe" ]