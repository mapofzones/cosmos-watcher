FROM golang:latest as build 

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 go build -o watcher ./cmd/watcher/main.go

FROM alpine:latest as production

WORKDIR /app

COPY --from=build /app/watcher  /app/watcher

CMD /app/watcher