package types

import (
	"io"
	"net/url"
	"os"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

// Site struct stores the site URL and sites that
// can be accessed through this Site
type Site struct {
	URL   *url.URL
	Links []*Site
	lock  sync.RWMutex
}

// AddLink adds links to Site
func (s *Site) AddLink(link *Site) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Links = append(s.Links, link)
}

// Print prints all the URLs in the site
// as well as in the sites linked to this site recursively
func (s *Site) Print(file *os.File, indent int) {
	a := strings.Join([]string{strings.Repeat("   ", indent), s.URL.String(), "\n"}, "")
	if file != nil {
		_, err := io.WriteString(file, a)
		if err != nil {
			return
		}
	} else {
		log.Println(a)
	}
	if len(s.Links) > 0 {
		d := strings.Join([]string{strings.Repeat("   ", indent+1), "Links:", "\n"}, "")
		if file != nil {
			_, err := io.WriteString(file, d)
			if err != nil {
				return
			}
		} else {
			log.Println(d)
		}
		for _, u := range s.Links {
			u.Print(file, indent+2)
		}
	}
}
