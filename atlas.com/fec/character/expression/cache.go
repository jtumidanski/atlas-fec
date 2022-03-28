package expression

import (
	"sync"
	"time"
)

type cache struct {
	mutex    sync.Mutex
	routines map[uint32]*Model
}

var c *cache
var once sync.Once

func getCache() *cache {
	once.Do(func() {
		c = &cache{
			mutex:    sync.Mutex{},
			routines: make(map[uint32]*Model, 0),
		}
	})
	return c
}

func (c2 *cache) add(characterId uint32, mapId uint32, expression uint32) *Model {
	expiration := time.Now().Add(time.Second * time.Duration(5))
	c2.mutex.Lock()
	e := &Model{
		characterId: characterId,
		mapId:       mapId,
		expression:  expression,
		expiration:  expiration,
	}
	c2.routines[characterId] = e
	c2.mutex.Unlock()
	return e
}

func (c2 *cache) popExpired() []*Model {
	es := make([]*Model, 0)
	c2.mutex.Lock()
	now := time.Now()
	for k, v := range c2.routines {
		if now.Sub(v.Expiration()) > 0 {
			es = append(es, v)
			delete(c2.routines, k)
		}
	}
	c2.mutex.Unlock()
	return es
}

func (c2 *cache) clear(characterId uint32) {
	c2.mutex.Lock()
	delete(c2.routines, characterId)
	c2.mutex.Unlock()
}
