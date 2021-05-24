package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hi20160616/fetchnews/config"
	"github.com/hi20160616/fetchnews/internal/pkg/service"
	"golang.org/x/sync/errgroup"
)

var (
	address string = config.Data.WebServer.Addr
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	// Web service
	s, err := service.NewServer(address)
	if err != nil {
		log.Println(err)
	}
	g.Go(func() error {
		defer cancel()
		return s.Start(ctx)
	})
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
		select {
		case sig := <-sigs:
			fmt.Println()
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			defer cancel()
			defer close(sigs)
			return s.Stop(ctx)
		case <-ctx.Done():
			defer cancel()
			defer close(sigs)
			return s.Stop(ctx)
		}
	})

	if err := g.Wait(); err != nil {
		log.Printf("main: %v", err)
	}
}
