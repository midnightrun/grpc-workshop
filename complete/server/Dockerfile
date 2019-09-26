FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build server.go

FROM scratch
COPY --from=builder /app/server /app/

ENTRYPOINT ["/app/server"]
