FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .
RUN go build -o ./bin/events-service ./cmd/main.go

FROM alpine as runner

COPY --from=builder build/bin/events-service /
CMD ["./events-service"]
