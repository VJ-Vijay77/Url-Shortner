package model

import (
	"time"
)

type Url struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	ShortUrl  string
	LongUrl   string
}
