package errors

import "errors"

// Error macros
var (
	ErrCrawlFailed          = errors.New("Crawling failed")
	ErrFileCreateFailed     = errors.New("Failed to create file")
	ErrURLfetchFailed       = errors.New("Failed to fetch urls")
	ErrURLparseFailed       = errors.New("Failed to parse URL")
	ErrHTMLfetchFailed      = errors.New("Failed to get HTML")
	ErrCreateDocumentFailed = errors.New("Failed to create document from HTML Body")
)
