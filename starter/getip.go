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
		Logger.Info("Error get checkip url")
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		Logger.Info("Error read checkip body")
		return nil, err
	}
	_ = resp.Body.Close()
	raw := strings.TrimSpace(string(b))
	return &raw, nil
}