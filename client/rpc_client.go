//+build ignore

package main

import (
	"context"
	"log"
	"os"
	"time"

	world_message "github.com/hellodudu/Ultimate/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:7080"
	defaultName = "nice to meet you"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := world_message.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &world_message.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	i := world_message.NewInviterClient(conn)
	nctx, ncancel := context.WithTimeout(context.Background(), time.Second)
	defer ncancel()
	ir, err := i.GetScore(nctx, &world_message.GetScoreRequest{Id: 1201616684167725058})
	if err != nil {
		log.Fatalf("could not invite: %v", err)
	}
	log.Printf("Inviter: %d", ir.Score)
}
