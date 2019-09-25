[![Build Status](https://travis-ci.org/rishubhjain/web-crawler.svg?branch=master)](https://travis-ci.org/rishubhjain/web-crawler)

# web-crawler
Web Crawler

Given an URL, it outputs textual sitemap, showing the links between pages. The crawler is limited to one subdomain
i.e when we start with *https://google.com/*, it would crawl all pages within google.com, but not follow external links, for example to facebook.com or community.google.com.


## Usage
```
$ make build && ./web-crawler  --hostURL https://google.com --depth 1| cat sitemap
```

## Testing
```
$ make test
```

## Build
```
$ make build
```

## Install
```
$ make install
$ web-crawler
```

## Uninstall
```
$ make uninstall
```
