package data

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/hi20160616/fetchnews-api/proto/v1"
	"github.com/hi20160616/fetchnews/config"
	"github.com/hi20160616/fetchnews/internal/pkg/db/ms"
)

func TestListArticles(t *testing.T) {
	if err := config.Reset("../../"); err != nil {
		t.Error(err)
	}
	if len(ms.Conns) == 0 {
		if err := ms.Open(); err != nil {
			t.Error(err)
		}
	}
	ar := NewArticleRepo(&Data{config.MicroService{Title: "bbc"}}, &log.Verbose{})
	as, err := ar.ListArticles(context.Background())
	if err != nil {
		t.Error(err)
	}
	for _, a := range as {
		fmt.Println(a)
	}
}

func TestGetArticle(t *testing.T) {
	if err := config.Reset("../../"); err != nil {
		t.Error(err)
	}
	if len(ms.Conns) == 0 {
		if err := ms.Open(); err != nil {
			t.Error(err)
		}
	}
	ar := NewArticleRepo(&Data{config.MicroService{Title: "bbc"}}, &log.Verbose{})
	as, err := ar.GetArticle(context.Background(), "cfb24f41b3786f04f821373233281d52")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(as.UpdateTime)
}

func TestSearchArticles(t *testing.T) {
	if err := config.Reset("../../"); err != nil {
		t.Error(err)
	}
	if len(ms.Conns) == 0 {
		if err := ms.Open(); err != nil {
			t.Error(err)
		}
	}
	ar := NewArticleRepo(&Data{config.MicroService{Title: "bbc"}}, &log.Verbose{})
	as, err := ar.SearchArticles(context.Background(), "奥运, 南海")
	if err != nil {
		t.Error(err)
	}
	for _, a := range as {
		fmt.Println(a.Title)
	}
}

func TestList(t *testing.T) {
	if err := config.Reset("../../"); err != nil {
		t.Error(err)
	}
	if len(ms.Conns) == 0 {
		if err := ms.Open(); err != nil {
			t.Error(err)
		}
	}
	as, err := List(context.Background(), &pb.ListArticlesRequest{})
	if err != nil {
		t.Error(err)
	}
	for _, a := range as {
		fmt.Println(a.Title, a.UpdateTime.AsTime())
	}
}
