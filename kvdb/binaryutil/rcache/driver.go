package rcache

import (
	"github.com/herb-go/datasource/kvdb"
)

type Driver interface {
	kvdb.Cache
	kvdb.Closer
	VersionStore
}
