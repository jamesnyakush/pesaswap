package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func (s Service) newReq(url string, body []byte, headers map[string]string) (string, error) {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))

	if err != nil {
		return "", err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(request)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return "", err
	}

	stringBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(stringBody), nil
}
func (s Service) getReq(url string, body []byte, headers map[string]string) (string, error) {
	request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(body))

	if err != nil {
		return "", err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(request)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return "", err
	}

	stringBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(stringBody), nil
}
