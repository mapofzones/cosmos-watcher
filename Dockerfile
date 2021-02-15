FROM alpine:latest as production

RUN apk add curl jq
WORKDIR /app

COPY ./watcher  /app/watcher
COPY scripts/run.sh /run.sh

CMD /run.sh