package server

import "DCache/HttpCache/cache"

var Cache cache.Cache

func init()  {
	Cache = cache.NewMemoryCache()
	_ = Cache.Set("houxinglin", []byte("dylan"))
}
