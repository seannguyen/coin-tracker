FROM coin-tracker-build-base

WORKDIR /go/src/github.com/seannguyen/coin-tracker
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["coin-tracker"]