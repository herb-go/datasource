package kvdb

import "time"

type Keyvalue interface {
	Set([]byte, []byte) error
	Get([]byte) ([]byte, error)
	Del([]byte) error
}

type Loader interface {
	Get([]byte) ([]byte, error)
}

type Cache interface {
	Get([]byte) ([]byte, error)
	SetWithTTL(key []byte, value []byte, ttl time.Duration) error
	Del([]byte) error
}
