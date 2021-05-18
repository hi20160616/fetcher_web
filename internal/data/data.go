package data

import (
	"context"
	"log"
	"time"

	"github.com/hi20160616/fetcher_web/config"
	pbBBC "github.com/hycka/news_fetcher/api/news_fetcher/api"
	"google.golang.org/grpc"
)

type Post struct {
	Id, Title, Content, WebSiteId, WebSiteTitle string
	UpdateTime                                  int64
}

func NewsSites() []config.Site {
	return config.Value.Sites
}

func List(ctx context.Context, domain string, id ...string) ([]*Post, error) {
	posts := []*Post{}
	switch domain {
	case "www.bbc.com":
		conn, err := grpc.Dial("localhost:10001", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pbBBC.NewNewsFetcherClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
		defer cancel()
		v, err := c.List(ctx, &pbBBC.ID{Id: "c263ded8cfbb5820e13e14b878680e90,3a7997ec004dad7d0befe251600d22e6"})
		if err != nil {
			log.Printf("c.GetPosts err: %+v", err)
		}
		for _, p := range v.Posts {
			posts = append(posts, &Post{
				Id:           p.Id,
				Title:        p.Title,
				Content:      p.Content,
				WebSiteId:    p.WebsiteId,
				WebSiteTitle: p.WebsiteTitle,
				UpdateTime:   p.UpdateTime,
			})
		}
	}
	return posts, nil
}
