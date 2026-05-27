package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
)

func main() {

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        panic(err)
    }

    server := grpc.NewServer()

    Register(server)

    log.Println("CFOS Scheduler Running")

    server.Serve(lis)
}