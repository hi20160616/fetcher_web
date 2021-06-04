package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/hi20160616/fetchnews-api/proto/v1"
)

type Data struct {
	MsTitle, ID, Keyword string
}

func ListArticles(ctx context.Context, in *pb.ListArticlesRequest, msTitle string) (*pb.ListArticlesResponse, error) {
	ar := NewArticleRepo(&Data{MsTitle: msTitle}, &log.Verbose{})
	as, err := ar.ListArticles(ctx)
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
	return &pb.ListArticlesResponse{Articles: resp}, nil
}

func GetArticle(ctx context.Context, in *pb.GetArticleRequest, msTitle string) (*pb.Article, error) {
	ar := NewArticleRepo(&Data{MsTitle: msTitle}, &log.Verbose{})
	a, err := ar.GetArticle(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Article{
		Id:            a.Id,
		Title:         a.Title,
		Content:       a.Content,
		WebsiteId:     a.WebsiteId,
		WebsiteDomain: a.WebsiteDomain,
		WebsiteTitle:  a.WebsiteTitle,
		UpdateTime:    a.UpdateTime,
	}, nil
}

func SearchArticles(ctx context.Context, in *pb.SearchArticlesRequest) (*pb.SearchArticlesResponse, error) {
	ar := NewArticleRepo(&Data{}, &log.Verbose{})
	as, err := ar.SearchArticles(ctx, in.Keyword)
	if err != nil {
		return nil, err
	}
	resp := []*pb.Article{}
	for _, a := range as {
		resp = append(resp, &pb.Article{
			Id:            a.Id,
			Title:         a.Title,
			Content:       a.Content,
			WebsiteId:     a.WebsiteId,
			WebsiteDomain: a.WebsiteDomain,
			WebsiteTitle:  a.WebsiteTitle,
			UpdateTime:    a.UpdateTime,
		})
	}
	return &pb.SearchArticlesResponse{Articles: resp}, nil
}
