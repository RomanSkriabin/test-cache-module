package cache

import (
	"errors"
	"time"
)

type Cache struct {
	items map[string]CacheItem
}

type CacheItem struct {
	value any
	ttl   time.Duration
}

var keyChan chan string

var keyTimer map[string]*time.Timer

func NewCache() Cache {
	keyChan = make(chan string)
	keyTimer = make(map[string]*time.Timer)
	return Cache{
		items: make(map[string]CacheItem),
	}
}

func (c *Cache) Get(key string) (value any, err error) {
	// получаем по ключу значение элемента

	cacheItem, ok := c.items[key]

	if ok {
		return cacheItem.value, nil
	}

	return "", errors.New("key not exist")
}

func (c *Cache) Delete(key string) (res bool, err error) {
	_, ok := c.items[key]

	if ok {
		delete(c.items, key)
		return true, nil
	}

	return false, errors.New("key not exist")
}

func (c *Cache) Set(key string, value any, ttl time.Duration) (res bool) {
	// проверим, есть ли такой ключ
	_, ok := c.items[key]
	if ok {
		// Удаление будет происходить по таймеру,
		// но если мы переопределили значение через set
		// необходимо стопнуть таймер, удалить и запустить заново
		keyTimer[key].Stop()
		delete(keyTimer, key)
	}
	keyTimer[key] = time.NewTimer(ttl)
	c.items[key] = CacheItem{value, ttl}
	go clearExpired(c, key, keyTimer[key])
	return true
}

func clearExpired(c *Cache, key string, t *time.Timer) bool {
	<-t.C
	c.Delete(key)
	return true
}
