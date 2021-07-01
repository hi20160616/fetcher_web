package ms

import (
	"context"
	"log"
	"time"

	pb "github.com/hi20160616/fetchnews-api/proto/v1"
	"github.com/pkg/errors"

	"github.com/hi20160616/fetchnews/configs"
	"google.golang.org/grpc"
)

type Conn struct {
	MicroService configs.MicroService
	ConnClient   *grpc.ClientConn
	FetchClient  pb.FetchNewsClient
}

var Conns = map[string]*Conn{}

func Open() error {
	if Conns != nil {
		for _, v := range configs.Data.MS {
			conn, err := grpc.Dial(v.Addr,
				grpc.WithInsecure(),
				grpc.WithBlock(),
				grpc.WithTimeout(5*time.Second),
				grpc.WithMaxMsgSize(32*10e6))
			if err != nil {
				return errors.WithMessagef(err, "[%s] grpc conn timeout", v.Title)
			}
			// defer conn.Close()
			Conns[v.Title] = &Conn{
				MicroService: v,
				ConnClient:   conn,
				FetchClient:  pb.NewFetchNewsClient(conn),
			}
		}

	}
	return nil
}

func Close() error {
	for _, c := range Conns {
		if err := c.ConnClient.Close(); err != nil {
			return err
		}
		delete(Conns, c.MicroService.Title)
	}
	return nil
}

// List will invoke microservice by title defined in config.json
func List(siteTitle string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := Conns[siteTitle].FetchClient.ListArticles(ctx, &pb.ListArticlesRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetArticles())
	return nil
}
