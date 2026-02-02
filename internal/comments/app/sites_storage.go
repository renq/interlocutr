package app

type SitesStorage interface {
	CreateSite(site Site) (Site, error)
	GetSite(ID string) (Site, error)
}
