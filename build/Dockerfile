FROM golang:alpine as builder
LABEL maintainer="mfathoor"

RUN apk update && apk add --no-cache git
RUN git config --global url."https://${GITHUB_USER}:${GITHUB_TOKEN}@${GITHUB_REPO}".insteadOf "https://${GITHUB_REPO}"

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o simkes ./cmd/web/main.go

FROM alpine:latest as release

WORKDIR /cmd
COPY --from=builder /app/simkes .

ENTRYPOINT ["/cmd/simkes"]
