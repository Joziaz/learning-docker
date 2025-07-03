FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY . .

RUN go build -o ./app.exe

FROM ubuntu 

WORKDIR /app

COPY --from=builder /go/src/app/app.exe ./

CMD ["./app.exe"]