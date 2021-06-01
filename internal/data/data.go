package data

import (
	"context"
	"log"
	"time"

	pb "github.com/hi20160616/fetchnews-api/proto/v1"
	"github.com/hi20160616/fetchnews/config"
	"google.golang.org/grpc"
)

type Data struct {
	MS config.MicroService
}

func NewsSites() map[string]config.MicroService {
	return config.Data.MS
}

func List(ctx context.Context, in *pb.ListArticlesRequest) ([]*pb.Article, error) {
	as := []*pb.Article{}
	conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFetchNewsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	resp, err := c.ListArticles(ctx, in)
	if err != nil {
		log.Printf("c.GetPosts err: %+v", err)
	}
	for _, item := range resp.Articles {
		as = append(as, item)
	}
	return as, nil
}
