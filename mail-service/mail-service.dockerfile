FROM alpine:latest

RUN mkdir /app

COPY bin/mailServiceApp /app
COPY templates /templates

CMD [ "/app/mailServiceApp" ]
