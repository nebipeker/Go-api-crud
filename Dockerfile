# Stage 1: Build the Go application
FROM golang:1.19 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Stage 2: Create the production image
FROM alpine:3.14

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=build /app/app .
COPY config.yaml .

CMD ["./app"]
