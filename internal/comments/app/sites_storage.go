package app

import "context"

type SitesStorage interface {
	CreateSite(ctx context.Context, site Site) (string, error) // TODO return ID only
	GetSite(ctx context.Context, ID string) (Site, error)
}
