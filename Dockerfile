FROM bitnami/golang:1.18-debian-11 as build

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 go build -o watcher ./cmd/watcher/main.go

FROM alpine:latest as production

RUN apk add curl jq coreutils
RUN apk add --no-cache bash
WORKDIR /app

COPY --from=build /app/watcher  /app/watcher
COPY scripts/run.sh /run.sh

CMD /run.sh