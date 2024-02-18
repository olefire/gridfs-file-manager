FROM alpine

COPY app.env .
ADD ./bin/app /app

RUN chmod +x ./app
ENTRYPOINT ["./app"]
