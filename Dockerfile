FROM golang:1.16

WORKDIR /go/src/app

COPY . .

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go build -o main .

EXPOSE 50051

CMD ["./main"]
