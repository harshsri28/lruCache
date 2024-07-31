package cache

import (
    "container/list"
    "sync"
    "time"

    "github.com/harshsri28/apica/helper"
)

type CacheItem struct {
    Key        string
    Value      string
    Expiration time.Time
    TimeExpiration time.Time
}

type LRUCache struct {
    capacity int
    items    map[string]*list.Element
    order    *list.List
    mutex    sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        items:    make(map[string]*list.Element),
        order:    list.New(),
    }
}

func (cache *LRUCache) Set(key string, value string, duration time.Duration) {
    cache.mutex.Lock()
    defer cache.mutex.Unlock()

    if element, found := cache.items[key]; found {
        cache.order.MoveToFront(element)
        item := element.Value.(*CacheItem)
        item.Value = value
        item.Expiration = time.Now().Add(duration)
        return
    }

    if cache.order.Len() >= cache.capacity {
        cache.evict()
    }

    item := &CacheItem{
        Key:        key,
        Value:      value,
        Expiration: time.Now().Add(duration),
        TimeExpiration: duration,
    }
    element := cache.order.PushFront(item)
    cache.items[key] = element
}

func (cache *LRUCache) Get(key string) (string, bool) {
    cache.mutex.Lock()
    defer cache.mutex.Unlock()

    if element, found := cache.items[key]; found {
        item := element.Value.(*CacheItem)
        if time.Now().After(item.Expiration) {
            cache.order.Remove(element)
            delete(cache.items, key)
            return "", false
        }
        element.Expiration = element.TimeExpiration
        cache.order.MoveToFront(element)
        return item.Value, true
    }

    return "", false
}

func (cache *LRUCache) Delete(key string) {
    cache.mutex.Lock()
    defer cache.mutex.Unlock()

    if element, found := cache.items[key]; found {
        cache.order.Remove(element)
        delete(cache.items, key)
    }
}

func (cache *LRUCache) evict() {
    element := cache.order.Back()
    if element != nil {
        cache.order.Remove(element)
        delete(cache.items, element.Value.(*CacheItem).Key)
    }
}

func (cache *LRUCache) StartExpirationRoutine() {
    go func() {
        for {
            time.Sleep(1 * time.Second)
            cache.mutex.Lock()
            for key, element := range cache.items {
                if time.Now().After(element.Value.(*CacheItem).Expiration) {
                    websocket.NotifyClients("expired", key)
                    cache.order.Remove(element)
                    delete(cache.items, key)
                }
            }
            cache.mutex.Unlock()
        }
    }()
}

func (cache *LRUCache) GetAll() map[string]string {
    cache.mutex.Lock()
    defer cache.mutex.Unlock()

    allItems := make(map[string]string)
    for key, element := range cache.items {
        item := element.Value.(*CacheItem)
        if time.Now().Before(item.Expiration) {
            allItems[key] = item.Value
        }
    }
    return allItems
}
