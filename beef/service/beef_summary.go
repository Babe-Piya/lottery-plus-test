package service

import (
	"io"
	"net/http"
	"strings"
)

type BeefSummaryResponse struct {
	Beef interface{} `json:"beef"`
}

var beefList = []string{"t-bone", "fatback", "pastrami", "pork", "meatloaf", "jowl", "enim", "bresaola"}

func (s service) BeefSummary() (BeefSummaryResponse, error) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return BeefSummaryResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	data := strings.ToLower(string(body))

	beefCount := make(map[string]int)
	for _, beef := range beefList {
		beefCount[beef] = strings.Count(data, beef)
	}

	return BeefSummaryResponse{
		Beef: beefCount,
	}, nil
}
