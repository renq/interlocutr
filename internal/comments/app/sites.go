package app

type Site struct {
	ID      string   `db:"id"`
	Domains []string `db:"domains"`
}

type CreateSiteRequest struct {
	ID      string   `json:"id"`
	Domains []string `json:"domains"`
}

type CreateSiteResponse struct {
	ID string `json:"id"`
}

type GetSiteRequest struct {
	ID string `param:"ID"`
}

type GetSiteResponse struct {
	ID      string   `json:"id"`
	Domains []string `json:"domains"`
}

func (a *App) CreateSite(command CreateSiteRequest) (CreateSiteResponse, error) {
	id, err := a.SitesStorage.CreateSite(Site(command))

	if err != nil {
		return CreateSiteResponse{}, err
	}

	return CreateSiteResponse{
		ID: id,
	}, nil
}

func (a *App) GetSite(command GetSiteRequest) (GetSiteResponse, error) {
	site, err := a.SitesStorage.GetSite(command.ID)

	if err != nil {
		return GetSiteResponse{}, err
	}

	return GetSiteResponse(site), nil
}
