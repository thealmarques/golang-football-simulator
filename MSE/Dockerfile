FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

COPY . /app

RUN go build

EXPOSE 9000

ENTRYPOINT [ "/app/MSE" ]