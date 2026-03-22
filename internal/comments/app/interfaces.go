package app

import "context"

type Transactor interface {
	Run(ctx context.Context, fn func(ctx context.Context) error) error
}

type SitesStorage interface {
	CreateSite(ctx context.Context, site Site) error
	GetSite(ctx context.Context, ID string) (Site, error)
}

type CommentsStorage interface {
	CreateComment(ctx context.Context, comment Comment) error
	GetComments(ctx context.Context, site, resource string) ([]Comment, error)
	Break()
}
