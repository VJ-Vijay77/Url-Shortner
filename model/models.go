package model

import (
	"time"
)

type Url struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	ShortUrl  string `gorm:"unique"`
	LongUrl   string `gorm:"unique"`
}

type UrlStruct struct {
	Url string `json:"url"`
}