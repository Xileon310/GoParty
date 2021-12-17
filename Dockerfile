FROM golang:1.17-alpine
LABEL maintainer="joseluis404@correo.ugr.es"

RUN addgroup -S goparty && adduser -S goparty -G goparty
USER goparty

WORKDIR /app/test
COPY go.mod ./

RUN go mod download
RUN go install github.com/go-task/task/v3/cmd/task@latest

ENTRYPOINT ["task", "test"]