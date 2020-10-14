package objectstore

import "time"

type Stat struct {
	Name         string
	IsFolder     bool
	Size         int64
	ModifiedTime time.Time
}
