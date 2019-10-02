package fetch

import (
	"context"
	"net/http"

	cerror "github.com/rishubhjain/web-crawler/errors"
	"github.com/rishubhjain/web-crawler/types"
	"github.com/rishubhjain/web-crawler/utils"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

// GoqueryFetcher structure implements Fetcher interface
type goqueryFetcher struct {
	client *http.Client
}

// NewGoqueryFetcher returns a crawler instance
func NewGoqueryFetcher(client *http.Client) Fetcher {
	return &goqueryFetcher{client: client}
}

// Fetch fetches URL using Goquery
func (h *goqueryFetcher) Fetch(ctx context.Context, site *types.Site) (err error) {

	// Local visited urls
	seenURLs := make(map[string]struct{})

	resp, err := h.client.Get(site.URL.String())
	if err != nil {
		log.WithFields(log.Fields{"Error": err,
			"URL": site.URL.String()}).Error(cerror.ErrGetRespFailed)
		return err
	}

	defer resp.Body.Close()
	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"Error": err,
			"URL": site.URL.String()}).Error(cerror.ErrCreateDocumentFailed)
		return err
	}

	// Fetch Links
	document.Find("a").Each(func(index int, element *goquery.Selection) {
		// Check whther the href attribute exists on the element or not
		href, exists := element.Attr("href")
		if exists {
			tempURL, err := utils.Parse(href)
			if err != nil {
				log.WithFields(log.Fields{"Error": err,
					"URL": href}).Debug("Failed to parse href")
				return
			}
			childURL := utils.ResolveURL(tempURL, site.URL)

			if childURL.Host != site.URL.Host {
				return
			}
			if _, ok := seenURLs[childURL.String()]; ok {
				return
			}

			// Storing the URL as seen URL
			seenURLs[childURL.String()] = struct{}{}

			tempSite := types.Site{URL: childURL}
			site.Links = append(site.Links, &tempSite)
		}
	})
	return nil

}
