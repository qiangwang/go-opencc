FROM golang:1.12.0

RUN apt-get update && apt-get install -y opencc libopencc-dev

WORKDIR /app
COPY . .