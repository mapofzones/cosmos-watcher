FROM golang:latest as build 
WORKDIR /app
COPY . /app
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build -o watcher ./cmd/watcher/main.go

FROM alpine:latest as production

RUN apk add curl jq
WORKDIR /app

COPY --from=build /app/watcher  /app/watcher
COPY scripts/run.sh /run.sh

CMD /run.sh