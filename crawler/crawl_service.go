package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type PdfFetcher struct {
	Client Client
}

func (p PdfFetcher) Fetch(number int) ([]byte, error) {
	r, err := p.Client.Get(fmt.Sprintf("https://localhost:9090/internet-bills/%d", number))
	if err != nil {

		return nil, err
	}
	return r, nil
}

type Client interface {
	Get(request string) ([]byte, error)
}

type HttpClientWrapper struct {
	client http.Client
}

func (h HttpClientWrapper) Get(request string) ([]byte, error) {
	resp, _ := h.client.Get(request)
	return ioutil.ReadAll(resp.Body)
}
