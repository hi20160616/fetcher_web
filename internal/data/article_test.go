package data

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/hi20160616/fetchnews-api/proto/v1"
	"github.com/hi20160616/fetchnews/configs"
	"github.com/hi20160616/fetchnews/internal/pkg/db/ms"
	"github.com/hi20160616/fetchnews/internal/service"
)

func TestListArticles(t *testing.T) {
	// path prepare
	if err := configs.Reset("../../"); err != nil {
		t.Error(err)
	}
	// init prepare
	if len(ms.Conns) == 0 {
		if err := ms.Open(); err != nil {
			t.Error(err)
		}
	}
	// test section
	ar := NewArticleRepo(&Data{MsTitle: "bbc"}, &log.Verbose{})
	as, err := ar.ListArticles(context.Background())
	if err != nil {
		t.Error(err)
	}
	for _, a := range as {
		fmt.Println(a)
	}
}

func TestGetArticle(t *testing.T) {
	// path prepare
	if err := configs.Reset("../../"); err != nil {
		t.Error(err)
	}
	// init prepare
	if len(ms.Conns) == 0 {
		if err := ms.Open(); err != nil {
			t.Error(err)
		}
	}
	// test section
	ar := NewArticleRepo(&Data{MsTitle: "bbc"}, &log.Verbose{})
	as, err := ar.GetArticle(context.Background(), "cfb24f41b3786f04f821373233281d52")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(as.UpdateTime)
}

func TestSearchArticles(t *testing.T) {
	// path prepare
	if err := configs.Reset("../../"); err != nil {
		t.Error(err)
	}
	// init prepare
	if len(ms.Conns) == 0 {
		if err := ms.Open(); err != nil {
			t.Error(err)
		}
	}
	// test section
	ar := NewArticleRepo(&Data{MsTitle: "bbc"}, &log.Verbose{})
	as, err := ar.SearchArticles(context.Background(), "奥运, 南海")
	if err != nil {
		t.Error(err)
	}
	for _, a := range as {
		fmt.Println(a.Title)
	}
}

func TestList(t *testing.T) {
	// path prepare
	if err := configs.Reset("../../"); err != nil {
		t.Error(err)
	}
	// init prepare
	if len(ms.Conns) == 0 {
		if err := ms.Open(); err != nil {
			t.Error(err)
		}
	}
	// test section
	as, err := service.ListArticles(context.Background(), &pb.ListArticlesRequest{}, "bbc")
	if err != nil {
		t.Error(err)
	}
	for _, a := range as.Articles {
		fmt.Println(a.Title, a.UpdateTime.AsTime())
	}
}
