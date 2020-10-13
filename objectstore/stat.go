package objectstore

import "time"

type Stat struct {
	Path         string
	IsFolder     bool
	Size         int
	ModifiedTime *time.Time
}
