package app

type SitesStorage interface {
	CreateSite(site Site) error
	GetSite(ID string) (Site, error)
}
