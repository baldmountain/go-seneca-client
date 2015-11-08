package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Requester is a struct to wrap a host and port to send requests to
type Requester struct {
	Host string
	Port int
}

func (r *Requester) request(json []byte) ([]byte, error) {
	url := fmt.Sprintf("http://%s:%d/act", r.Host, r.Port)
	res, err := http.Post(url, "text/json", bytes.NewReader(json))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	return b, err
}

// Close does nothing
func (r *Requester) Close() error {
	return nil
}

// Act on an interface to request and fills in a responce interface
func (r *Requester) Act(req interface{}, res interface{}) error {
	s, err := json.Marshal(req)
	if err != nil {
		return err
	}

	out, err := r.request(s)
	if err != nil {
		return err
	}

	return json.Unmarshal(out, res)
}
