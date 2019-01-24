FROM golang:1.11.5-alpine as builder
COPY . .
RUN GO111MODULE=on go build -v -o main .

FROM alpine:latest
COPY --from=builder main /usr/local/bin/clipper
