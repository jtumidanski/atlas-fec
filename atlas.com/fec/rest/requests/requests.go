package requests

import (
	"atlas-fec/rest/attributes"
	"net/http"
)

const (
	baseRequest string = "http://atlas-nginx:80"
)

func get(url string, resp interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}

	err = processResponse(r, resp)
	return err
}

func processResponse(r *http.Response, rb interface{}) error {
	err := attributes.FromJSON(rb, r.Body)
	if err != nil {
		return err
	}

	return nil
}

