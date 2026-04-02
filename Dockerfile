FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/Nutrimama cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates curl tzdata && update-ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/Nutrimama .

ARG APPPORT=3014
EXPOSE ${APPPORT}

ENTRYPOINT [ "./Nutrimama" ]