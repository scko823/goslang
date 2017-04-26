FROM golang

WORKDIR /app

COPY *.go /app/

RUN rm db.go

RUN go get github.com/gorilla/websocket

RUN go get github.com/go-redis/redis

RUN mkdir assets/

COPY assets/vendor/ /app/assets/vendor/

VOLUME /app/assets

EXPOSE 8080

RUN go build -o goslang

CMD ["./goslang"]
