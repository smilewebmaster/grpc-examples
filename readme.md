### Under this repository, I will share my gRPC examples and the languages used will be GO, JS

#### GO proto file compile
    protoc -I . reverse.proto --go_out=plugins=grpc:.
