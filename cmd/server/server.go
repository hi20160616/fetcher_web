package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hi20160616/fetchnews/config"
	"github.com/hi20160616/fetchnews/internal/server"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

var (
	address string = config.Data.WebServer.Addr
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// Web server
	s, err := server.NewServer(address)
	if err != nil {
		log.Println(err)
	}
	g.Go(func() error {
		log.Println("Server start on " + address)
		return s.Start(ctx)
	})
	g.Go(func() error {
		<-ctx.Done() // wait for stop signal
		log.Println("Server stop now...")
		return s.Stop(ctx)
	})

	// Elegant stop
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	g.Go(func() error {
		select {
		case sig := <-sigs:
			fmt.Println()
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			log.Printf("not canceled by context: %s", err)
		} else {
			log.Println(err)
		}
	}
}
