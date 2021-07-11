package starter

import (
	"io"
	"net/http"
	"strings"
)

const infoUrl = "https://checkip.amazonaws.com"

func GetLocalIp() (*string, error) {
	resp, err := http.Get(infoUrl)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	_ = resp.Body.Close()
	raw := strings.TrimSpace(string(b))
	return &raw, nil
}