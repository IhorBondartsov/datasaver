FROM golang:1.11.2-alpine3.8 AS builder

RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/IhorBondartsov/datasaver

ENV GO111MODULE=on


COPY go.mod .
COPY go.sum .
RUN go mod download


# This image builds the weavaite server
FROM builder AS server_builder
# Here we copy the rest of the source code
COPY . .
# And compile the project
RUN ls
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"'
#In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine AS server
# We add the certificates to be able to verify remote weaviate instances
RUN apk add ca-certificates
# Finally we copy the statically compiled Go binary.
COPY --from=server_builder /go/bin/datasaver /bin/datasaver
EXPOSE 1812
ENTRYPOINT ["/bin/datasaver"]
 