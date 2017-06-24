FROM scko823/goslang-ui:latest

FROM golang

WORKDIR /app

COPY *.go /app/

RUN rm db.go

RUN go get github.com/gorilla/websocket

RUN go get github.com/go-redis/redis

COPY --from=0 /usr/src/app/assets/ assets/

EXPOSE 8080

RUN go build -o goslang

CMD ["./goslang"]
