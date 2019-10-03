package crawler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rishubhjain/web-crawler/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCrawl(t *testing.T) {

	body := tests.ParseHTML(t, "../tests/fixtures/test.html")
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// Send response to be tested
		rw.Write([]byte(body))
	}))
	// Close the server when test finishes
	defer server.Close()

	mockParser := new(tests.HTTPParserMock)
	mockParser.On("Parse", mock.Anything).Return(nil)
	crawler := NewCrawler()
	baseURL := server.URL
	depth := 1
	site, _ := crawler.Crawl(baseURL, depth)
	assert.Equal(t, len(site.Links), 2)

	baseURL = "h//*goo%&^gle"
	_, err := crawler.Crawl(baseURL, depth)
	assert.NotNil(t, err)

}
