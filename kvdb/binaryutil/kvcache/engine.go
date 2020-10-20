package kvcache

import (
	"time"
)

type Engine struct {
	versionGenerator func() (string, error)
	versionstore     VersionStore
	driver           Driver
}

func (e *Engine) NewCache(path []byte, irrevocable bool) *Cache {
	return &Cache{
		Path:        path,
		irrevocable: irrevocable,
		engine:      e,
	}
}

var DefaultVersionbGenerator = func() (string, error) {
	buf := make([]byte, 8)
	DataOrder.PutUint64(buf, uint64(time.Now().UnixNano()))
	return string(buf), nil
}

func NewEngine() *Engine {
	return &Engine{
		versionGenerator: DefaultVersionbGenerator,
	}
}
