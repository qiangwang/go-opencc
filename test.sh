#/bin/bash

docker build -t go-opencc .
docker run --rm go-opencc bash -c "go build && go test"