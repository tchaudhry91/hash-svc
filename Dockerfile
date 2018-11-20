# Builder 
FROM     golang:latest as BUILDER
COPY    . /go/src/github.com/tchaudhry91/hash-svc
WORKDIR /go/src/github.com/tchaudhry91/hash-svc
RUN     go get -d -v ./... && \
        go test -v ./...
RUN     cd cmd && go build -o hash-svc


# Final Image
FROM       alpine
RUN        apk update && apk add --no-cache ca-certificates
COPY       --from=BUILDER /go/src/github.com/tchaudhry91/hash-svc/cmd/hash-svc /bin/
ENTRYPOINT [ "/bin/hash-svc" ]
