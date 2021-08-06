FROM golang:1.16 as build

WORKDIR /app

COPY . /app

#RUN apk add --no-cache gcc musl-dev
RUN apt-get update && apt-get install -y make gcc gawk bison libc-dev
#RUN apt-get add --no-cache make gcc gawk bison linux-headers libc-dev
RUN go build -o watcher ./cmd/watcher/main.go
#RUN CGO_ENABLED=0 go build -o watcher ./cmd/watcher/main.go
RUN ls -la /app

FROM alpine:latest as production

RUN apk add curl jq
RUN apk add --no-cache bash
WORKDIR /app

COPY --from=build /app/watcher  /app/watcher
COPY scripts/run.sh /run.sh

RUN ls -la /
RUN ls -la /app/
CMD /run.sh
