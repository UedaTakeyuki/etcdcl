package main

import (
	"context"
	"strconv"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func BenchmarkPut(b *testing.B) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cli.Put(context.Background(), "hello", strconv.Itoa(i))
	}
}

func BenchmarkGet(b *testing.B) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cli.Get(context.Background(), "hello")
	}
}
