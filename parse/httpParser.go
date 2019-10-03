package parse

import (
	"context"
	"net/http"

	"github.com/rishubhjain/web-crawler/types"
	"github.com/rishubhjain/web-crawler/utils"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

// httpParser structure implements Parser interface
type httpParser struct {
	client *http.Client
}

// NewHTTPParser returns a parser instance
func NewHTTPParser(client *http.Client) Parser {
	return &httpParser{client: client}
}

// Parse extracts URL from web page using html tokenizer
func (h *httpParser) Parse(ctx context.Context, site *types.Site) (err error) {

	// Local visited urls
	seenRefs := make(map[string]struct{})

	rootURL := site.URL

	retry := types.RetryingClient{
		Client:      h.client,
		MaxAttempts: 2,
	}

	// Get Page from URL
	resp, err := retry.Get(rootURL.String())
	if err != nil {
		// Logged in the caller function
		return err
	}

	tokenizer := html.NewTokenizer(resp.Body)
	defer resp.Body.Close()

	for {
		// get the next token type
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			return
		}

		if tokenType != html.StartTagToken {
			continue
		}

		token := tokenizer.Token()
		// Extract links
		if token.DataAtom.String() != "a" && token.DataAtom.String() != "link" {
			continue
		}

		for _, attr := range token.Attr {
			if attr.Key != "href" {
				continue
			}
			tempURL, err := utils.Parse(attr.Val)
			if err != nil {
				// Debug logs
				log.WithFields(log.Fields{"Error": err,
					"URL": attr.Val}).Debug("Failed to parse href")
				continue
			}
			childURL := utils.ResolveURL(tempURL, site.URL)
			if childURL.Host != site.URL.Host {
				continue
			}
			if _, ok := seenRefs[childURL.String()]; ok {
				continue
			}

			// Storing the URL as seen URL
			seenRefs[childURL.String()] = struct{}{}
			tempSite := types.Site{URL: childURL}
			site.Links = append(site.Links, &tempSite)
		}
	}
}
