package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func Test_run(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("error: failed to listen port: %+v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx, l)
	})
	url := fmt.Sprintf("http://%s", l.Addr())
	rsp, err := http.Get(url)
	if err != nil {
		t.Errorf("error: faield to get: %+v", err)
	}
	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Errorf("error: faied to read response body: %+v", err)
	}
	if string(got) != "hello" {
		t.Errorf("error: unexpected response body: expected hello, but got %s", string(got))
	}
	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
