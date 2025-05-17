package cache

import (
	"sync"

	"github.com/google/uuid"
	"github.com/okarpova/my-app/internal/domain"
)

type Cache struct {
	mx sync.RWMutex
	m  map[uuid.UUID]domain.Profile
}

func New() *Cache {
	return &Cache{
		m: make(map[uuid.UUID]domain.Profile),
	}
}

func (c *Cache) Add(key uuid.UUID, profile domain.Profile) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.m[key] = profile
}

func (c *Cache) Get(key uuid.UUID) (domain.Profile, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	p, ok := c.m[key]
	if !ok {
		return p, domain.ErrNotFound
	}

	return p, nil
}

func (c *Cache) Update(key uuid.UUID, profile domain.Profile) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	_, ok := c.m[key]
	if !ok {
		return domain.ErrNotFound
	}
	c.m[key] = profile
	return nil
}

func (c *Cache) Delete(key uuid.UUID) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	_, ok := c.m[key]
	if !ok {
		return domain.ErrNotFound
	}

	delete(c.m, key)
	return nil
}
