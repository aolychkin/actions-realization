package sqlite

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

// Конструктор Storage
func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	// Указываем путь до файла БД
	db, err := gorm.Open(sqlite.Open(storagePath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
