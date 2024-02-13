FROM golang:alpine
LABEL maintainer="mfathoor.23@gmail.com" \
      name="mfathoor/simkes-api" \
      github="https://github.com/fathoor/simkes-api" \
      dockerhub="https://hub.docker.com/r/mfathoor/simkes-api"

RUN apk update && apk add --no-cache git

WORKDIR /cmd

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o simkes-api

EXPOSE ${APP_PORT}

ENTRYPOINT ["/cmd/simkes-api"]
