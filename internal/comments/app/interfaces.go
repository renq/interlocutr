package app

import "context"

type SitesStorage interface {
	CreateSite(ctx context.Context, site Site) (string, error)
	GetSite(ctx context.Context, ID string) (Site, error)
}

type CommentsStorage interface {
	CreateComment(ctx context.Context, comment Comment) error
	GetComments(ctx context.Context, site, resource string) ([]Comment, error)
}
