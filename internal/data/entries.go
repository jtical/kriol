//Filename: internal/data/entries.go

package data

import (
	"time"
)

// holds entries informatiom
// back tick character(struct tags) shows how the key should be formated
type Entries struct {
	ID        int64     `json:"id"`
	English   string    `json:"english"`
	Kriol     string    `json:"kriol,mitemoty"`
	CreatedAt time.Time `json:"created_at"`
	Version   int32     `json:"version"`
}
