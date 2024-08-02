package api

import (
	"encoding/json"
	"net/http"
	"time"
)

func httpFetch(url string, target interface{}) error {
	var httpclient = &http.Client{Timeout: 10 * time.Second}

	res, err := httpclient.Get(url)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}
