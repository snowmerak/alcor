FROM alpine:latest

RUN apk add libc6-compat
COPY ./bin/alcor-linux-amd64 ./

ENTRYPOINT ./alcor-linux-amd64