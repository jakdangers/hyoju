package internal

import "github.com/google/uuid"

type Model struct {
	UUID uuid.UUID `gorm:"primaryKey"`
}
