package handler

import (
	"io"
	"net/http"
	"strings"
)

var (
	ZoneIdList map[string]string // zoneDomain -> zoneId
	ZoneList   map[string]string // zoneId -> zoneDomain
)

func Init() {
	ZoneIdList = make(map[string]string)
}

func sendReq(method string, path string, header map[string]string, body string) (string, error) {
	// get a httpclient
	client := &http.Client{}
	url := "https://api.cloudflare.com/client/v4" + path
	bodyReader := strings.NewReader(body)
	req, _ := http.NewRequest(method, url, bodyReader)
	for k, v := range header {
		req.Header.Set(k, v)
	}

	if resp, err := client.Do(req); err != nil {
		return "", err
	} else {
		buf := new(strings.Builder)
		io.Copy(buf, resp.Body)
		return buf.String(), nil
	}
}
