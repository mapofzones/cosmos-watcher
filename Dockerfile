FROM bitnami/golang:1.20-debian-11 as build

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . /app

RUN apt-get update && apt-get install -y make gcc gawk bison libc-dev

RUN go build -o watcher ./cmd/watcher/main.go

RUN ls -la /app

FROM ubuntu:latest as production

RUN apt-get update && apt-get install -y curl jq coreutils dos2unix

COPY --from=build /app/libwasmvm.so /usr/lib/libwasmvm.so
COPY --from=build /app/libwasmvm.x86_64.so /usr/lib/libwasmvm.x86_64.so
COPY --from=build /app/watcher  /app/watcher
COPY --from=build /app/scripts/run.sh /run.sh

RUN dos2unix /run.sh
CMD ["/run.sh"]