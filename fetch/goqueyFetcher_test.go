package fetch

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/rishubhjain/web-crawler/tests"
	"github.com/rishubhjain/web-crawler/types"
	"github.com/rishubhjain/web-crawler/utils"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {

	body := tests.ParseHTML(t, "../tests/fixtures/google.com.html")
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// Send response to be tested
		rw.Write([]byte(body))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	client := NewGoqueryFetcher(server.Client())

	URL, err := utils.Parse(server.URL)
	assert.Nil(t, err)
	site := types.Site{URL: URL, Links: nil}

	err = client.Fetch(context.Background(), &site)
	assert.Nil(t, err)
	assert.Equal(t, len(site.Links), 2)

	site = types.Site{URL: &url.URL{
		Scheme: "http",
		Host:   "www.google",
	},
		Links: nil}

	err = client.Fetch(context.Background(), &site)
	assert.NotNil(t, err)

}
