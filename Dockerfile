FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

EXPOSE 8080

COPY . .

RUN go mod tidy
RUN go build -o binary cmd/app/main.go

ENTRYPOINT ["/app/binary"]