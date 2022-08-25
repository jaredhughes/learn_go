FROM alpine:latest

RUN mkdir /app

COPY bin/frontEndApp /app

CMD ["/app/frontEndApp"]
