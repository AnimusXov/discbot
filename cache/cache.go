package cache

import (
	"github.com/akyoto/cache"
	"time"
)

var c = cache.New(30 * time.Minute)

func Put[T any](k T, v T) {
	c.Set(k, v, 0)
}

func Read[T any](k T) (obj interface{}, found bool) {
	return c.Get(k)
}
