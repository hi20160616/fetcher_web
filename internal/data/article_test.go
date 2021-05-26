package data

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hi20160616/fetchnews/config"
)

func TestListArticles(t *testing.T) {
	ar := NewArticleRepo(&Data{config.MicroService{Title: "bbc"}}, &log.Verbose{})
	as, err := ar.ListArticles(context.Background())
	if err != nil {
		t.Error(err)
	}
	for _, a := range as {
		fmt.Println(a)
	}
}
