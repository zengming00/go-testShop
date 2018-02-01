package lib

import (
	"sync"
	"time"
)

type CacheMgr struct {
	mu            sync.RWMutex
	gcIntervalSec int64
	items         map[string]*CacheItem
}

type CacheItem struct {
	value         interface{}
	expireTimeSec int64
}

func NewCacheMgr(gcIntervalSec int64) *CacheMgr {
	mgr := &CacheMgr{
		gcIntervalSec: gcIntervalSec,
		items:         make(map[string]*CacheItem),
	}
	go mgr.gc()
	return mgr
}

func (mgr *CacheMgr) gc() {
	for {
		<-time.After(time.Duration(mgr.gcIntervalSec) * time.Second)
		mgr.mu.Lock()
		for key := range mgr.items {
			mgr.isExpired(key)
		}
		mgr.mu.Unlock()
	}
}

func (mgr *CacheMgr) isExpired(key string) bool {
	if item, ok := mgr.items[key]; ok {
		if item.expireTimeSec <= 0 {
			return false
		}
		if item.expireTimeSec < time.Now().Unix() {
			delete(mgr.items, key)
			return true
		}
		return false
	}
	return true
}

func (mgr *CacheMgr) Add(key string, v int64, expireSec int64) int64 {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	expireTimeSec := int64(-1)
	if expireSec > 0 {
		expireTimeSec = time.Now().Unix() + expireSec
	}

	if item, ok := mgr.items[key]; ok {
		if oldVal, ok := item.value.(int64); ok {
			item.expireTimeSec = expireTimeSec
			item.value = oldVal + v
			return oldVal
		}
	}
	mgr.items[key] = &CacheItem{
		value:         v,
		expireTimeSec: expireTimeSec,
	}
	return 0
}

func (mgr *CacheMgr) Set(key string, value interface{}, expireSec int64) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	expireTimeSec := int64(-1)
	if expireSec > 0 {
		expireTimeSec = time.Now().Unix() + expireSec
	}
	mgr.items[key] = &CacheItem{
		value:         value,
		expireTimeSec: expireTimeSec,
	}
}

func (mgr *CacheMgr) Get(key string) (interface{}, bool) {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()
	if !mgr.isExpired(key) {
		if v, ok := mgr.items[key]; ok {
			return v.value, true
		}
	}
	return nil, false
}

func (mgr *CacheMgr) Del(key string) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	delete(mgr.items, key)
}

func (mgr *CacheMgr) Flush() {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()
	mgr.items = make(map[string]*CacheItem)
}
