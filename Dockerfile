FROM golang:1.11.5 as builder
COPY . /src
WORKDIR /src
RUN GO112MODULE=on go build -ldflags "-linkmode external -extldflags -static" -o main -a main.go


FROM alpine:latest
COPY --from=builder /src/main /usr/local/bin/clipper
