package service

type service struct {
}

type Service interface {
	BeefSummary() (BeefSummaryResponse, error)
}

func New() Service {
	return service{}
}
