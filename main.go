package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	c "github.com/rishubhjain/web-crawler/crawler"
	cerror "github.com/rishubhjain/web-crawler/errors"

	log "github.com/sirupsen/logrus"
)

var (
	hostURL = flag.String("hostURL", "https://google.com",
		"host url to crawl")
	depth = flag.Int("depth", 1, "depth")
)

func main() {
	flag.Parse()
	now := time.Now()

	site, err := c.NewCrawler().Crawl(*hostURL, *depth)
	if err != nil {
		log.WithFields(log.Fields{"Error": err,
			"URL": hostURL}).Error(cerror.ErrCrawlFailed)
		return
	}
	elapsed := time.Since(now)

	file, err := os.Create("sitemap")
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).
			Error(cerror.ErrFileCreateFailed)
		return
	}

	defer file.Close()
	// Print Sitemap in file
	// if nil is passed then print the links in lognil
	site.Print(file, 0)

	fmt.Println(elapsed)
}
