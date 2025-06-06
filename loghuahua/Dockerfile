FROM golang:1.24-alpine AS builder

WORKDIR /app
# COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download

COPY main.go ./
RUN CGO_ENABLED=0 go build -o loghuahua main.go

FROM scratch
WORKDIR /
COPY --from=builder /app/loghuahua .
ENTRYPOINT ["/loghuahua"]