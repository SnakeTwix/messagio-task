FROM golang:1.22.5-alpine as base
WORKDIR /app


FROM base as dev

RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download

CMD ["air"]