# STEP 1 build executable binary
FROM golang:alpine AS builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

ENV GO111MODULE=on
WORKDIR /go/src/hello-go

COPY . .

RUN go get -d -v ./...
RUN go build -v


# STEP 2 build a small image
FROM alpine
LABEL maintainer="Sir Rob <sirrob@netsak.de>"

WORKDIR /app
COPY --from=builder /go/src/hello-go/hello-go /app/

ENTRYPOINT ["/app/hello-go"]
