package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/hi20160616/fetchnews-api/proto/v1"
	"github.com/hi20160616/fetchnews/internal/biz"
	"github.com/hi20160616/fetchnews/internal/pkg/db/ms"
)

var _ biz.ArticleRepo = new(articleRepo)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper("articleRepo", logger),
	}
}

func (ar *articleRepo) ListArticles(ctx context.Context) ([]*biz.Article, error) {
	as := []*biz.Article{}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	c := ms.Conns[ar.data.MS.Title].FetchClient
	resp, err := c.ListArticles(ctx, &pb.ListArticlesRequest{})
	if err != nil {
		return nil, err
	}
	for _, item := range resp.Articles {
		t := &biz.Article{
			Id:            item.Id,
			Title:         item.Title,
			Content:       item.Content,
			WebsiteId:     item.WebsiteId,
			WebsiteDomain: item.WebsiteDomain,
			WebsiteTitle:  item.WebsiteTitle,
		}
		as = append(as, t)
	}
	return as, nil
}

func (ar *articleRepo) GetArticle(ctx context.Context, id string) (*biz.Article, error) {
	return nil, nil
}

func (ar *articleRepo) SearchArticles(ctx context.Context, keyword ...string) ([]*biz.Article, error) {
	return nil, nil
}
