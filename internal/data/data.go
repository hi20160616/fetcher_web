package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/hi20160616/fetchnews-api/proto/v1"
	"github.com/hi20160616/fetchnews/config"
)

type Data struct {
	MS config.MicroService
}

func List(ctx context.Context, in *pb.ListArticlesRequest) ([]*pb.Article, error) {
	ar := NewArticleRepo(&Data{config.MicroService{Title: "bbc"}}, &log.Verbose{})
	as, err := ar.ListArticles(context.Background())
	if err != nil {
		return nil, err
	}
	resp := []*pb.Article{}
	for _, a := range as {
		resp = append(resp, &pb.Article{
			Id:            a.Id,
			Title:         a.Title,
			Content:       a.Content,
			UpdateTime:    a.UpdateTime,
			WebsiteId:     a.WebsiteId,
			WebsiteDomain: a.WebsiteDomain,
			WebsiteTitle:  a.WebsiteTitle,
		})
	}
	return resp, nil
}
