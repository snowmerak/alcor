package csset

import "sync"

type CSSet struct {
	mutex sync.RWMutex
	table map[string]struct{}
}

func (c *CSSet) Put(data string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, ok := c.table[data]; ok {
		return false
	}
	c.table[data] = struct{}{}
	return true
}

func (c *CSSet) Exist(data string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if _, ok := c.table[data]; ok {
		return true
	}
	return false
}

func (c *CSSet) Remove(data string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.table, data)
}
