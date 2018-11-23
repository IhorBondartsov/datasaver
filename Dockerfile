FROM golang:1.11.2-alpine3.8 AS server_builder

RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/IhorBondartsov/datasaver

ENV GO111MODULE=on
COPY . .

RUN \
    go mod download && \
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"'

FROM alpine AS server

RUN apk add ca-certificates
COPY --from=server_builder /go/bin/datasaver /bin/datasaver
EXPOSE 1812
ENTRYPOINT ["/bin/datasaver"]
 