package cache

import "sync"

type MemoryCache struct {
	c			map[string][]byte
	rwMutex		sync.RWMutex
	Stat
}

func (m *MemoryCache) Set(k string, v []byte) error {
	m.rwMutex.Lock()
	defer m.rwMutex.Unlock()
	// 如果k存在，更新keysize和valsize
	if val, exist := m.c[k]; exist {
		m.del(k, val)
	}
	m.c[k] = v
	m.add(k, v)
	return nil
}

func (m *MemoryCache) Get(k string) ([]byte, error) {
	m.rwMutex.RLock()
	defer m.rwMutex.RUnlock()
	return m.c[k], nil
}

func (m *MemoryCache) Del(k string) error {
	m.rwMutex.Lock()
	defer m.rwMutex.Unlock()
	if v, exist := m.c[k]; exist {
		delete(m.c, k)
		m.del(k, v)
	}
	return nil
}

func (m *MemoryCache) GetStat() Stat {
	return m.Stat
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		c:       make(map[string][]byte),
		rwMutex: sync.RWMutex{},
		Stat:    Stat{},
	}
}
