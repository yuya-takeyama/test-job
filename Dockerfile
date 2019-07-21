FROM golang:1.12.7-alpine3.9 as builder

RUN apk --update add gcc

WORKDIR /app
COPY . /app

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"'

FROM alpine:3.9

COPY --from=builder /app/test-job /usr/local/bin

ENTRYPOINT ["test-job"]
