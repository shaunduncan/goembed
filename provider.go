package goembed

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	urllib "net/url"
)

type Provider struct {
	Name     string
	Endpoint *urllib.URL
}

func (p Provider) makeRequestUrl(url string) string {
	// Make a copy of the endpoint
	u, _ := urllib.Parse(p.Endpoint.String())
	values := urllib.Values{
		"url":    []string{url},
		"format": []string{"xml"},
	}
	u.RawQuery = values.Encode()

	return u.String()
}

func (p Provider) get(url string) (Resource, error) {
	resp, err := http.Get(p.makeRequestUrl(url))
	if err != nil {
		return Resource{}, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Resource{}, err
	}

	resource := new(Resource)

	if err := xml.Unmarshal(contents, resource); err != nil {
		return Resource{}, err
	}

	return *resource, nil
}

func (p Provider) OEmbed(urls ...string) []Resource {
	resChan := make(chan Resource)
	resources := []Resource{}

	for _, url := range urls {
		go func(u string, c chan Resource) {
			res, _ := p.get(url)
			c <- res
		}(url, resChan)
	}

	for i := 0; i < len(urls); i++ {
		resources = append(resources, <-resChan)
	}

	return resources
}
