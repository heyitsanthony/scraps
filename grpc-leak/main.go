package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
)

func get(kvc pb.KVClient) {
	reqrange := &pb.RangeRequest{Key: []byte("mykey")}
	if _, err := kvc.Range(context.TODO(), reqrange, grpc.FailFast(false)); err != nil {
		panic(err)
	}
}

func testGet(num int) {
	conn, err := grpc.Dial("127.0.0.1:2379", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	kvc := pb.NewKVClient(conn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			get(kvc)
		}()
	}

	wg.Wait()
	fmt.Println("Done.")
}

func main() {
	testGet(50000)
	time.Sleep(60 * time.Minute)
}
