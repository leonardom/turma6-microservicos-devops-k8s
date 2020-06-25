FROM golang:1.14 AS builder

ADD src src/

WORKDIR /go/src/server

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -extldflags "-static"' -o /go/bin/server

FROM scratch

COPY --from=builder /go/bin/server /server

ENTRYPOINT ["./server"]