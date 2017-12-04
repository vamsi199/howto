package main

import (
	"fmt"
	pb "github.com/muly/howto/golang/web/grpc/echo-example/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

func main() {

	fmt.Println("listening...")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("could not listen:",err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type server struct{}

func (server) EchoEcho(ctx context.Context, i *pb.Input) (*pb.Output, error) {
	o:= pb.Output{i.Text}
	return &o, nil
}
