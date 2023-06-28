package repositories

import "gorm.io/gorm"

type PGRepository struct {
	DB *gorm.DB
}

func NewPGDatabase() *PGRepository {
	db, _ := GormInit()
	return &PGRepository{DB: db}
}
