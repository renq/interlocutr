package app

type SitesStorage interface {
	CreateSite(site Site) (string, error) // TODO return ID only
	GetSite(ID string) (Site, error)
}
