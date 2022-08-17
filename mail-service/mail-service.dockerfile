FROM alpine:latest

RUN mkdir /app

COPY bin/mailServiceApp /app

CMD [ "/app/mailServiceApp" ]
