package cache

import "errors"

type Cache struct {
	items map[string]any
}

func NewCache() Cache {
	return Cache{
		items: make(map[string]any),
	}
}

func (c Cache) get(key string) (value any, err error) {
	// получаем по ключу значение элемента

	value, ok := c.items[key]

	if ok == true {
		return value, nil
	}

	return "", errors.New("Key not exist")
}

func (c Cache) delete(key string) (res bool, err error) {
	_, ok := c.items[key]

	if ok == true {
		delete(c.items, key)
		return true, nil
	}

	return false, errors.New("Key not exist")
}

func (c Cache) set(key string, value any) (res bool) {
	c.items[key] = value
	return true
}
