import (
	"sync"
	"time"
)

type CacheManager struct {
	lock   *sync.RWMutex
	caches map[string]*Cache
}

type Cache struct {
	Value interface{}
	Time  int64
}

func NewCacheManager(size int) *CacheManager {
	return &CacheManager{new(sync.RWMutex), make(map[string]*Cache, size)}
}

func (this *CacheManager) Set(key string, v *Cache) {
	this.lock.Lock()
	this.caches[key] = v
	this.lock.Unlock()
}

func (this *CacheManager) Get(key string) *Cache {
	this.lock.Lock()
	v := this.caches[key]
	this.lock.Unlock()
	return v
}

func (this *CacheManager) Delete(key string) (v *Cache) {
	this.lock.Lock()
	v = this.caches[key]
	delete(this.caches, key)
	this.lock.Unlock()
	return v
}

func (this *CacheManager) IsExpired(key string, ttl int) bool {
	if v := this.Get(key); v != nil {
		return (time.Now().Unix() - v.Time) >= int64(ttl)
	} else {
		return true
	}
}