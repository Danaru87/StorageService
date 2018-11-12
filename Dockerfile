FROM golang:1.11
MAINTAINER Charly Caulet <contact@charly-caulet.net>

WORKDIR /app
COPY ./src .
RUN go get
RUN go build

CMD ["StorageService"]