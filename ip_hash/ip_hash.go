package iphash

import (
	"errors"
	"net/url"
	"sync"
	"unsafe"

	"github.com/cespare/xxhash"
	roundrobin "github.com/ductnn/lba/round_robin"
)

// ErrServersNotExist is the error that server does not exist
var ErrServersNotExist = errors.New("server does not exist")

// Interface
type IPHash interface {
	Next(in *url.URL) *url.URL
}

type iphash struct {
	urls []*url.URL
	cnt  uint64
	m    map[uint64]*url.URL
	mu   *sync.RWMutex
	rr   roundrobin.RoundRobin
}

// New returns IPHash(*iphash) object
func New(urls []*url.URL) (IPHash, error) {
	if len(urls) == 0 {
		return nil, ErrServersNotExist
	}

	rr, _ := roundrobin.New(urls)

	return &iphash{
		urls: urls,
		cnt:  uint64(len(urls)),
		m:    make(map[uint64]*url.URL),
		mu:   new(sync.RWMutex),
		rr:   rr,
	}, nil
}

func (i *iphash) Next(in *url.URL) *url.URL {
	hashN := xxhash.Sum64(*(*[]byte)(unsafe.Pointer(&in.Host))) % i.cnt

	i.mu.RLock()
	if url, ok := i.m[hashN]; ok {
		i.mu.RUnlock()
		return url
	}
	i.mu.RUnlock()

	url := i.rr.Next()
	i.mu.Lock()
	i.m[hashN] = url
	i.mu.Unlock()

	return url
}
