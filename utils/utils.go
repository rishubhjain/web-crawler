package utils

import "net/url"

// Parse returns parsed URL
func Parse(hostURL string) (*url.URL, error) {
	finalURL, err := url.Parse(hostURL)
	return finalURL, err
}

// HasSameHost finds out whether two URLs have same host or not
func HasSameHost(hostURL, siteURL *url.URL) bool {
	if hostURL.Host == siteURL.Host {
		return true
	}
	return false
}

// ResolveURL returns resolved URL
func ResolveURL(hostURL, siteURL *url.URL) *url.URL {
	newURL := siteURL.ResolveReference(hostURL)
	newURL.Fragment = ""
	return newURL
}
