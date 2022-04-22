package tool

import (
	"io/ioutil"
	"net/http"
)

// GetIPFromAPI gets the IP address from the API
func GetIPFromAPI() (string, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.ipify.org?format=raw", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
