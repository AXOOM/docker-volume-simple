FROM golang:1.9-alpine as builder
RUN apk add --no-cache gcc libc-dev
WORKDIR /go/src/github.com/axoom/docker-volume-simple
COPY . .
RUN go install --ldflags '-extldflags "-static"'

FROM alpine as rootfs
RUN mkdir -p /run/docker/plugins /mnt
COPY --from=builder /go/bin/docker-volume-simple .

FROM docker
WORKDIR /plugin
COPY config.json .
COPY --from=rootfs / rootfs/
