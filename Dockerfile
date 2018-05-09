# build binary
FROM seannguyen/coin-tracker-build-base as build-base

WORKDIR /go//src/github.com/seannguyen/coin-tracker
COPY . .

RUN go get -d -v ./...
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -extldflags "-static"' -o coin-tracker main.go

# package binary
FROM alpine:3.7

RUN apk add --update ca-certificates
COPY --from=build-base /go/src/github.com/seannguyen/coin-tracker/coin-tracker /var/app/coin-tracker

ENTRYPOINT ["/var/app/coin-tracker"]