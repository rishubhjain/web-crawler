package types

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func populateSite(URL string) *Site {

	site := Site{URL: &url.URL{
		Scheme: "http",
		Host:   URL,
	},
		Links: nil}
	return &site
}

func TestAddLink(t *testing.T) {
	s := populateSite("www.google.com")
	s.AddLink(populateSite("www.facebook.com"))
	assert.Equal(t, len(s.Links), 1)
}

func TestPrint(t *testing.T) {
	s := populateSite("www.google.com")
	s.Print(nil, 0)
	file, err := os.Create("testFile")
	if err != nil {
		return
	}
	defer file.Close()
	s.Print(file, 0)

	newSite := Site{URL: &url.URL{
		Scheme: "http",
		Host:   "www.facebook.com",
	},
		Links: nil}
	s.Links = append(s.Links, &newSite)
	s.Print(file, 0)
}
