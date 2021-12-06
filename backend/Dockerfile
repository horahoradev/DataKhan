FROM golang:1.16.5-buster as builder

WORKDIR /datakhan/backend

# build binary
COPY . /datakhan/backend

RUN go mod vendor && \
 go build -mod=vendor -o /backend.bin

ENTRYPOINT ["/backend.bin"]