package data

import "github.com/hi20160616/fetcher_web/config"

func NewsSites() []config.Site {
	return config.Value.Sites
}
