package parse

import (
	"context"

	"github.com/rishubhjain/web-crawler/types"
)

// Parser abstracts Parse functionality
type Parser interface {
	Parse(ctx context.Context, site *types.Site) (err error)
}
