FROM bitnami/golang:1.21-debian-11 as build 

ENV GOPROXY=https://proxy.golang.org

WORKDIR /app

COPY . /app

RUN go build -tags 'secretcli' -o watcher ./cmd/watcher/main.go

FROM ubuntu:latest as production

RUN apt-get update && apt-get install -y curl jq coreutils dos2unix
WORKDIR /app

COPY --from=build /app/watcher  /app/watcher
COPY --from=build /app/scripts/run.sh /run.sh

RUN dos2unix /run.sh
CMD ["/run.sh"]
