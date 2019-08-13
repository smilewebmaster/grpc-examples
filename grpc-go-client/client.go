package main

import (
    "context"
    "fmt"
    "os"
    pb "github.com/batuhangoksu/grpc-examples/grpc-go-client/reverse"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
)

func main() {
    opts := []grpc.DialOption{
        grpc.WithInsecure(),
    }
    args := os.Args
    conn, err := grpc.Dial("127.0.0.1:5300", opts...)
    if err != nil {
        grpclog.Fatalf("fail to dial: %v", err)
    }
    defer conn.Close()
    client := pb.NewReverseClient(conn)
    request := &pb.Request{
        Message: args[1],
    }
    response, err := client.Do(context.Background(), request)
    if err != nil {
        grpclog.Fatalf("fail to dial: %v", err)
    }
   fmt.Println(response.Message)
}