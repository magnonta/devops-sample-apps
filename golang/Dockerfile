FROM golang:1.20-alpine AS build

WORKDIR /app

COPY . .

ENV GOARCH=amd64
ENV GOOS=linux

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/* .

EXPOSE 8080

CMD ["./main"]
