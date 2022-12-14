package leastconnections

import (
	"net/url"
	"sync"

	"github.com/pkg/errors"
)

// ErrServersNotExist is the error that server does not exist
var ErrServersNotExist = errors.New("server does not exist")

// Interface
type LeastConnections interface {
	Next() (next *url.URL, done func())
}

type conn struct {
	url *url.URL
	cnt int
}

type leastConnections struct {
	conns []conn
	mu    *sync.Mutex
}

// New initializes a new instance of LeastConnected
func New(urls []*url.URL) (LeastConnections, error) {
	if len(urls) == 0 {
		return nil, ErrServersNotExist
	}

	conns := make([]conn, len(urls))
	for i := range conns {
		conns[i] = conn{
			url: urls[i],
			cnt: 0,
		}
	}

	return &leastConnections{
		conns: conns,
		mu:    new(sync.Mutex),
	}, nil
}

func (lc *leastConnections) Next() (*url.URL, func()) {
	var (
		min = -1
		idx int
	)

	lc.mu.Lock()

	for i, conn := range lc.conns {
		if min == -1 || conn.cnt < min {
			min = conn.cnt
			idx = i
		}
	}
	lc.conns[idx].cnt++

	lc.mu.Unlock()

	var done bool
	return lc.conns[idx].url, func() {
		lc.mu.Lock()
		if !done {
			lc.conns[idx].cnt--
			done = true
		}
		lc.mu.Unlock()
	}
}
