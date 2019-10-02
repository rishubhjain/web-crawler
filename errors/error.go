package errors

import "errors"

// Error macros
var (
	ErrCrawlFailed          = errors.New("Crawling failed")
	ErrFileCreateFailed     = errors.New("Failed to create file")
	ErrFetchFailed          = errors.New("Failed to fetch site from url")
	ErrURLparseFailed       = errors.New("Failed to parse URL")
	ErrGetRespFailed        = errors.New("Failed to get response")
	ErrCreateDocumentFailed = errors.New("Failed to create document from HTML Body")
)
