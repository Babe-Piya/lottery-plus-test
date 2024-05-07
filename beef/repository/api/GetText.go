package api

import (
	"io"
	"strings"
)

func (a api) GetText() (res string, err error) {
	resp, err := a.Client.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	data := strings.ToLower(string(body))

	return data, nil
}
