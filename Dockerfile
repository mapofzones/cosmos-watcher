FROM bitnami/golang:1.18-debian-11 as build

WORKDIR /app

COPY . /app

RUN go build -o watcher ./cmd/watcher/main.go

FROM ubuntu:latest as production

RUN apt-get update && apt-get install -y curl jq coreutils

COPY --from=build /app/watcher  /app/watcher
COPY --from=build /app/scripts/run.sh /run.sh

CMD ["/run.sh"]