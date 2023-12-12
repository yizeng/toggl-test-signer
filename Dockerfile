FROM golang:1.21 AS base

FROM base AS development

WORKDIR /project

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
