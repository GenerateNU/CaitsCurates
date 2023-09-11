FROM golang:1.21.1


RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main ./src/main.go

EXPOSE 8080
CMD [ "/app/main" ]