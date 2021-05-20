package data

import (
	"context"
	"log"
	"time"

	pb "github.com/hi20160616/fetchnews/api/fetchnews/v1"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func ListArticles(ctx context.Context, in *pb.ListArticlesRequest) (*pb.Articles, error) {
	as := &pb.Articles{}
	conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
		return nil, errors.WithMessage(err, "ListArticles Dial error")
	}
	defer conn.Close()
	c := pb.NewNewsFetcherClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	resp, err := c.ListArticles(ctx, in)
	if err != nil {
		log.Printf("c.GetPosts err: %+v", err)
		return nil, errors.WithMessage(err, "ListArticles invoke error")
	}
	for _, item := range resp.Articles.Articles {
		as.Articles = append(as.Articles, item)
	}
	return as, nil
}
