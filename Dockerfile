FROM golang:1.25-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify

EXPOSE 8080

CMD ["./app"]
