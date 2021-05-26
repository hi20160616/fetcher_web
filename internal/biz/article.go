package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Article struct {
	Id, Title, Content, WebsiteId, WebsiteDomain, WebsiteTitle string
	UpdateTime                                                 timestamppb.Timestamp
}

type ArticleRepo interface {
	ListArticles(ctx context.Context) ([]*Article, error)
	GetArticle(ctx context.Context, id string) (*Article, error)
	SearchArticles(ctx context.Context, keyword ...string) ([]*Article, error)
}

type ArticleUsecase struct {
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

func (uc *ArticleUsecase) List(ctx context.Context) ([]*Article, error) {
	return uc.repo.ListArticles(ctx)
}

func (uc *ArticleUsecase) Get(ctx context.Context, id string) (*Article, error) {
	return uc.repo.GetArticle(ctx, id)
}

func (uc *ArticleUsecase) Search(ctx context.Context, keyword ...string) ([]*Article, error) {
	return uc.repo.SearchArticles(ctx, keyword...)
}
