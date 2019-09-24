package utils

import "net/url"

// Parse returns parsed URL
func Parse(rawurl string) (*url.URL, error) {
	targetURL, err := url.Parse(rawurl)
	return targetURL, err
}

// HasSameHost finds out whether two URLs have same host or not
func HasSameHost(rawURL, siteURL *url.URL) bool {
	if rawURL.Host == siteURL.Host {
		return true
	}
	return false
}

// ResolveURL returns resolved URL
func ResolveURL(rawURL, siteURL *url.URL) *url.URL {
	newURL := siteURL.ResolveReference(rawURL)
	newURL.Fragment = ""
	return newURL
}
