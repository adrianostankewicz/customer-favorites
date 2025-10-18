FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY . .

#RUN go test -v -count=1 ./...

RUN go build -o app ./cmd/customer/main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /build/app .

EXPOSE 3000

CMD ["./app"]