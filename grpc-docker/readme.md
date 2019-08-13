# gRPC Docker Demo

### Docker cmd
    docker build -t grpc-docker-demo .
    docker run -it -p 5300:5300 grpc-docker-demo:latest (or container name)

### Proto file compile
    protoc -I . reverse.proto --go_out=plugins=grpc:.