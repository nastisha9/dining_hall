FROM golang:1.19.1-bullseye

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

EXPOSE 3030
CMD ["/app/main"]