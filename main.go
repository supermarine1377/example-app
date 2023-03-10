package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("fatal: application port number not found as environment variable")
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Printf("error: failed to listen port: %+v", err)
	}
	if err := run(context.Background(), l); err != nil {
		log.Printf("error: failed to run server; %+v", err)
	}
}

func run(ctx context.Context, l net.Listener) error {
	s := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "hello")
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Printf("error: failed to close server: %+v\n", err)
			return err
		}
		return nil
	})
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("error: failed to shut down server: %+v\n", err)
	}
	return eg.Wait()
}
