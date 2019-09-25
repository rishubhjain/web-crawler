package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/rishubhjain/web-crawler/web"

	log "github.com/sirupsen/logrus"
)

var (
	hostURL = flag.String("hostURL", "http://google.com", "host url to crawl")
	depth   = flag.Int("depth", 1, "depth")
)

func main() {
	flag.Parse()
	now := time.Now()

	site, err := web.NewCrawler().Crawl(*hostURL, *depth)
	if err != nil {
		log.WithFields(log.Fields{"Error": err, "URL": hostURL}).Error("Failed to crawl")
		return
	}
	elapsed := time.Since(now)

	file, err := os.Create("sitemap")
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("Failed to open file")
		return
	}
	defer file.Close()
	site.Print(file, 0)
	fmt.Println(elapsed)
}
