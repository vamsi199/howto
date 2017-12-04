package main

import (
	"context"
	pb "github.com/muly/howto/golang/web/grpc/echo-example/pb"
	"google.golang.org/grpc"
	"fmt"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("dial error: ",err)
		return
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)

	text := &pb.Input{Text: "hello, world"}
	resp, err := client.EchoEcho(context.Background(), text)
	if err != nil {
		fmt.Println("client call error: ",err)
		return
	}
	fmt.Println("response is:",resp.Text)

}
