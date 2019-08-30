FROM golang:1.12.9-alpine3.10 AS builder

COPY . /go/src/github.com/subpathdev/kubeedge-database

RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/kubeedge-database -ldflags="-w -s" github.com/subpathdev/kubeedge-database

FROM alpine:3.10

COPY --from=builder /usr/local/bin/kubeedge-database kubeedge-database

ENTRYPOINT ["/kubeedge-database"]
