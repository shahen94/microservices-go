FROM alpine:latest

WORKDIR /app

COPY broker-service .

CMD [ "./broker-service" ]