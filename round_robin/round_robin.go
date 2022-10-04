package roundrobin

import (
	"errors"
	"net/url"
	"sync/atomic"
)

// ErrServersNotExist is the error that server does not exist
var ErrServersNotExist = errors.New("server does not exist")

// Interface
type RoundRobin interface {
	Next() *url.URL
}

type roundrobin struct {
	urls []*url.URL
	next uint32
}

// New returns RoundRobin implementation.
func New(urls []*url.URL) (RoundRobin, error) {
	if len(urls) == 0 {
		return nil, ErrServersNotExist
	}

	return &roundrobin{
		urls: urls,
	}, nil
}

// Next returns next server
func (r *roundrobin) Next() *url.URL {
	n := atomic.AddUint32(&r.next, 1)
	return r.urls[(int(n)-1)%len(r.urls)]
}
