package types

import "sync"

// Set stores a hashmap for URL, used for storing visited URL
type Set struct {
	URL map[string]bool
	// to synchronize access to the set
	lock sync.RWMutex
}

// Add adds a new url to the URL Set. Returns a pointer to the Set.
func (s *Set) Add(newURL string) *Set {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.URL == nil {
		s.URL = make(map[string]bool)
	}
	_, ok := s.URL[newURL]
	if !ok {
		s.URL[newURL] = true
	}
	return s
}

// Delete removes the string from the Set and returns Has(string)
func (s *Set) Delete(newurl string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.URL[newurl]
	if ok {
		delete(s.URL, newurl)
	}
	return ok
}

// Has returns true if the Set contains the string
func (s *Set) Has(newurl string) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.URL[newurl]
	return ok
}

// Size returns the size of the set
func (s *Set) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.URL)
}
