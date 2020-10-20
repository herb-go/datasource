package kvcache

import "time"

type Driver interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, data []byte, ttl time.Duration) error
	Del(key []byte) error
	Open() error
	Close() error
	VersionStore
}
