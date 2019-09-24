package fetch

import (
	"context"
	"net/http"

	"github.com/rishubhjain/web-crawler/types"
	"github.com/rishubhjain/web-crawler/utils"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

// HTTPFetcher structure implements Fetcher interface
type httpFetcher struct {
	client *http.Client
}

// NewHTTPFetcher returns a crawler instance
func NewHTTPFetcher(client *http.Client) Fetcher {
	return &httpFetcher{client: client}
}

// Fetch fetches URL using tokenizer
func (h *httpFetcher) Fetch(ctx context.Context, site *types.Site) (err error) {

	// Local visited urls
	seenRefs := make(map[string]struct{})

	rootURL := site.URL
	resp, err := h.client.Get(rootURL.String())
	if err != nil {
		log.WithFields(log.Fields{"Error": err,
			"URL": rootURL.String()}).Error("Failed to get HTML")
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
		// Get links
		if token.DataAtom.String() != "a" && token.DataAtom.String() != "link" {
			continue
		}

		for _, attr := range token.Attr {
			if attr.Key != "href" {
				continue
			}
			tempURL, err := utils.Parse(attr.Val)
			if err != nil {
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
