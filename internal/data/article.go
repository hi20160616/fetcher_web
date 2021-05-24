package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hi20160616/fetchnews/internal/biz"
)

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
	return nil, nil
}

func (ar *articleRepo) GetArticle(ctx context.Context, id string) (*biz.Article, error) {
	return nil, nil
}

func (ar *articleRepo) SearchArticles(ctx context.Context, keyword ...string) ([]*biz.Article, error) {
	return nil, nil
}
