#! /bin/bash
rm -rf assets/
unzip -d . assets.zip
CGO_ENABLED=0 GOOS=linux /usr/local/go/bin/go build -a -installsuffix cgo -o main .
docker build -t goslang:latest -f Dockerfile.scratch .