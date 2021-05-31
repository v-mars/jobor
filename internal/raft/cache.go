package raft

import (
	"encoding/json"
	"io"
	"sync"
)

type cacheManager struct {
	data map[string]string
	sync.RWMutex
}

func NewCacheManager() *cacheManager {
	cm := &cacheManager{}
	cm.data = make(map[string]string)
	return cm
}

func (c *cacheManager) Get(key string) string {
	c.RLock()
	ret := c.data[key]
	c.RUnlock()
	return ret
}

func (c *cacheManager) Delete(key string) string {
	c.RLock()
	delete (c.data, key)
	c.RUnlock()
	return ""
}

func (c *cacheManager) Set(key string, value string) error {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
	return nil
}

// Marshal serializes cache data
func (c *cacheManager) Marshal() ([]byte, error) {
	c.RLock()
	defer c.RUnlock()
	dataBytes, err := json.Marshal(c.data)
	return dataBytes, err
}

// UnMarshal deserializes cache data
func (c *cacheManager) UnMarshal(serialized io.ReadCloser) error {
	var newData map[string]string
	if err := json.NewDecoder(serialized).Decode(&newData); err != nil {
		return err
	}

	c.Lock()
	defer c.Unlock()
	c.data = newData

	return nil
}
