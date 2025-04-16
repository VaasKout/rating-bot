# syntax=docker/dockerfile:1
FROM golang:1.24 AS rating-bot

WORKDIR /app
COPY rating-bot ./
RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o rating-bot ./cmd/main.go

FROM alpine:3.18
RUN apk add --update --no-cache openssh sshpass bash rsync
WORKDIR /go/src/rating-bot/
RUN ls
COPY --from=rating-bot /app ./
CMD ["./rating-bot"]
