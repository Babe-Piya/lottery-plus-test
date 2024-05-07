package service

import "lottery-plus/beef/repository/api"

type service struct {
	API api.API
}

type Service interface {
	BeefSummary() (BeefSummaryResponse, error)
}

func New(api api.API) Service {
	return service{
		API: api,
	}
}
