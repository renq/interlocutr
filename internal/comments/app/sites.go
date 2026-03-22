package app

import "context"

type Site struct {
	ID      string
	Domains []string
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

func (a *App) CreateSite(ctx context.Context, command CreateSiteRequest) (CreateSiteResponse, error) {
	err := a.SitesStorage.CreateSite(ctx, Site(command))

	if err != nil {
		return CreateSiteResponse{}, err
	}

	return CreateSiteResponse{
		ID: command.ID,
	}, nil
}

func (a *App) GetSite(ctx context.Context, command GetSiteRequest) (GetSiteResponse, error) {
	site, err := a.SitesStorage.GetSite(ctx, command.ID)

	if err != nil {
		return GetSiteResponse{}, err
	}

	return GetSiteResponse(site), nil
}
