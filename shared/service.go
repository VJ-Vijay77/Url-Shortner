package shared

import "gorm.io/gorm"


type Service struct {
	Db *gorm.DB
}

func NewUserService(db *gorm.DB) *Service {
	return &Service{Db: db}
}