# syntax=docker/dockerfile:1

# parent image
FROM golang:1.18-alpine AS builder

# Build Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

#install build dependencies
RUN apk add build-base

# workspace directory
WORKDIR /app

# copy `go.mod` and `go.sum`
ADD go.mod go.sum ./

# install dependencies
RUN go mod download

# copy source code
COPY . .

# build executable
RUN go build -gcflags="all=-N -l" -o build/ ./cmd/storxscan

##################################

# parent image
FROM alpine:3.12.2

# copy binary file from the `builder` stage
COPY --from=builder /app/build/storxscan /var/lib/storx/go/bin/
COPY --from=builder /go/bin/dlv /var/lib/storx/go/bin/
ADD entrypoint.sh /var/lib/storx/entrypoint.sh

ENV PATH="/var/lib/storx/go/bin/:${PATH}"

# exec
ENTRYPOINT ["/var/lib/storx/entrypoint.sh"]