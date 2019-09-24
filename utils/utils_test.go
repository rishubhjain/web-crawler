package utils

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	targetURL, err := Parse("https://google.com")
	assert.Nil(t, err)
	assert.Equal(t, targetURL.Host, "google.com")

	_, err = Parse("h//*goo%&^gle")
	assert.NotNil(t, err)
}

func TestHasSameHost(t *testing.T) {
	testURL1, _ := Parse("https://www.google.com")
	testURL2, _ := Parse("https://www.google.com/doodles")
	assert.True(t, HasSameHost(testURL1, testURL2))
	testURL3, _ := Parse("https://www.facebook.com")
	assert.False(t, HasSameHost(testURL1, testURL3))
}

func TestResolveUrl(t *testing.T) {
	testURL1 := &url.URL{
		Scheme: "http",
		Host:   "www.google.com",
	}

	testURL2, _ := Parse("https://www.google.com/doodles")
	resolvedURL := ResolveURL(testURL1, testURL2)
	assert.Equal(t, "http://www.google.com", resolvedURL.String())

}
