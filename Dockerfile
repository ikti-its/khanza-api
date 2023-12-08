FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o simkes-api

EXPOSE ${APP_PORT}

ENTRYPOINT ["/app/simkes-api"]
