FROM golang:1.21-alpine

WORKDIR /app
ADD . /app

RUN go build -o /gateway-service

EXPOSE 10000

CMD [ "/gateway-service" ]