package goembed

import (
	"encoding/json"
	"encoding/xml"
)

type Author struct {
	Name string
	URL  string
}

type Resource struct {
	// For xml encoding
	XMLName xml.Name `json:"-" xml:"oembed"`

	Type            string `json:"type" xml:"type"`
	Version         string `json:"version" xml:"version"`
	Title           string `json:"title,omitempty" xml:"title,omitempty"`
	AuthorName      string `json:"author_name,omitempty" xml:"author_name,omitempty"`
	AuthorURL       string `json:"author_url,omitempty" xml:"author_url,omitempty"`
	ProviderName    string `json:"provider_name,omitempty" xml:"provider_name,omitempty"`
	ProviderURL     string `json:"provider_url,omitempty" xml:"provider_url,omitempty"`
	CacheAge        int    `json:"cache_age,omitempty" xml:"cache_age,omitempty"`
	ThumbnailURL    string `json:"thumbnail_url,omitempty" xml:"thumbnail_url,omitempty"`
	ThumbnailWidth  int    `json:"thumbnail_width,omitempty" xml:"thumbnail_width,omitempty"`
	ThumbnailHeight int    `json:"thumbnail_height,omitempty" xml:"thumbnail_height,omitempty"`

	// Photo, Video, Rich Types
	Width  int `json:"width,omitempty" xml:"width,omitempty"`
	Height int `json:"height,omitempty" xml:"height,omitempty"`

	// Photo Types
	URL string `json:"url,omitempty" xml:"url,omitempty"`

	// Video, Rich Types
	HTML string `json:"html,omitempty" xml:"html,omitempty"`

	// Raw data
	Raw string `json:"-" xml:"-"`
}

func (r Resource) XML() ([]byte, error) {
	return xml.Marshal(r)
}

func (r Resource) JSON() ([]byte, error) {
	return json.Marshal(r)
}
