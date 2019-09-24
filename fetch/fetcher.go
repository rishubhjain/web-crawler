package fetch

import (
	"context"

	"github.com/rishubhjain/web-crawler/types"
)

// Fetcher abstracts fetching functionality
type Fetcher interface {
	Fetch(ctx context.Context, site *types.Site) (err error)
}
