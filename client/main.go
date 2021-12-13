package main

import (
	"context"
	"flag"
	"log"
  "sync"
	"time"
	"google.golang.org/grpc"
  pb "github.com/ritmatter/gocounter/counter"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
  ngoroutine = flag.Int("n", 10, "how many goroutines")
)

func Increment(ctx context.Context, amount int64, c pb.CounterClient, wg *sync.WaitGroup) {
  defer wg.Done()

	r, err := c.Increment(ctx, &pb.IncrementRequest{Amount: amount})
	if err != nil {
		log.Fatalf("could not increment: %v", err)
	}
	log.Printf("Received: %d after input amount %d", r.GetNewTotal(), amount)
}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCounterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 300 * time.Second)
	defer cancel()

  var wg sync.WaitGroup
  for i := 0; i < *ngoroutine; i++ {
    wg.Add(1)
    go Increment(ctx, int64(i), c, &wg)
  }

  wg.Wait()
}
