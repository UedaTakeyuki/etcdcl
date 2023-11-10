package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()

	//var resp *clientv3.PutResponse
	//var result interface{}
	for i := 0; i < 1000; i++ {
		//resp, err = put(cli, i)
		err = get(cli)
	}
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	} else {
		//log.Println(resp)
		//log.Println(result)
	}

}

func get(cli *clientv3.Client) ( /*value interface{}, */ err error) {
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	//var resp *clientv3.GetResponse
	_, err = cli.Get(ctx, "hello")
	//	value = resp.Node.Value
	return
}

func put(cli *clientv3.Client, n int) (resp *clientv3.PutResponse, err error) {
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	resp, err = cli.Put(ctx, "hello", strconv.Itoa(n))

	return
}
