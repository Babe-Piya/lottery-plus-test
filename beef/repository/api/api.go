package api

//go:generate mockgen -source=api.go -destination=mocks/mock_api.go

import "net/http"

type api struct {
	Client http.Client
}

type API interface {
	GetText() (res string, err error)
}

func New(client http.Client) API {
	return api{
		Client: client,
	}
}
