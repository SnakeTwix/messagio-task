FROM golang:1.22.5-alpine AS base
WORKDIR /app

FROM base AS prod

COPY go.mod go.sum ./
RUN go mod download

COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin ./cmd/messagio/main.go

CMD ["./bin/main"]


FROM base AS dev

RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download

CMD ["air"]