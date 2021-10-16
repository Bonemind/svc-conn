# Build
FROM golang:alpine AS build

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN go build -o /svc-conn

# Run
FROM alpine:latest

WORKDIR /app

COPY --from=build /svc-conn ./svc-conn
COPY templates/ ./templates/

EXPOSE 7887

ENTRYPOINT ["/app/svc-conn"]
