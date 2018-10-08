FROM golang:1.11 as builder
WORKDIR /go/src/github.com/lvrach/mem-eater
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mem-eater ./cmd/mem-eater/main.go

FROM scratch
COPY --from=builder /go/src/github.com/lvrach/mem-eater .
CMD ["./mem-eater"]
