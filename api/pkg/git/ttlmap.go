package git

import (
	"sync"
	"time"
)

type item struct {
	lastAccess int64
}

type TTLMap struct {
	m map[string]*item
	l sync.Mutex
}

func NewTTLMap(ln int, maxTTL int) (m *TTLMap) {
	m = &TTLMap{m: make(map[string]*item, ln)}
	go func() {
		for now := range time.Tick(time.Minute) {
			m.l.Lock()
			for k, v := range m.m {
				if now.Unix()-v.lastAccess > int64(maxTTL) {
					delete(m.m, k)
				}
			}
			m.l.Unlock()
		}
	}()
	return
}

func (m *TTLMap) Put(k string) {
	m.l.Lock()
	defer m.l.Unlock()

	it, ok := m.m[k]
	if !ok {
		it = &item{}
		m.m[k] = it
	}
	it.lastAccess = time.Now().Unix()
}

func (m *TTLMap) Ok(k string) bool {
	m.l.Lock()
	defer m.l.Unlock()
	if it, ok := m.m[k]; ok {
		it.lastAccess = time.Now().Unix()
		return true
	}

	return false
}

func (m *TTLMap) Delete(k string) {
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.m, k)
}
