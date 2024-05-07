package service

import (
	"strings"
)

type BeefSummaryResponse struct {
	Beef interface{} `json:"beef"`
}

var beefList = []string{"t-bone", "fatback", "pastrami", "pork", "meatloaf", "jowl", "enim", "bresaola"}

func (s service) BeefSummary() (BeefSummaryResponse, error) {
	data, err := s.API.GetText()
	if err != nil {
		return BeefSummaryResponse{}, err
	}

	beefCount := make(map[string]int)
	for _, beef := range beefList {
		beefCount[beef] = strings.Count(data, beef)
	}

	return BeefSummaryResponse{
		Beef: beefCount,
	}, nil
}
