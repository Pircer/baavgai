FROM alpine

WORKDIR /app

COPY bin/baavgai .
COPY bin/config.yaml .

EXPOSE 8000
ENTRYPOINT [". /baavgai"]

