FROM alpine:latest

WORKDIR /app

COPY auth-service .

CMD [ "./auth-service" ]