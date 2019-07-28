# STEP 1 build executable binary
FROM golang:alpine AS builder

WORKDIR /go/src/hello
COPY . .

RUN go get -d -v ./...
RUN go build -v


# STEP 2 build a small image
FROM alpine
LABEL maintainer="Sir Rob <sirrob@netsak.de>"

WORKDIR /app
COPY --from=builder /go/src/hello/hello /app/


ENTRYPOINT ["/app/hello"]

