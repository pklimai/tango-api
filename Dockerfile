# Golang image on apline linux with Golang version corresponding your go.mod
FROM golang:1.22 AS build

# Env variables
ENV GOOS linux
ENV CGO_ENABLED 0

# Work directory
WORKDIR /app

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copying all the files
COPY . .

# Building the application
RUN go build -o bin/tango-api cmd/tango-api/main.go

# Fetching the latest nginx image
FROM alpine:3.20 AS prod

# Work directory
WORKDIR /app

# Copying built assets from build
COPY --from=build app/bin/tango-api .
COPY --from=build app/config/.env config/

# Starting our application
CMD ./tango-api

# Exposing server port
# HTTP, Swagger, gRPC
EXPOSE 7000 7001 7002
