package models

import (
	"time"
)

// User représente le modèle d'une sonnerie
type Ringtone struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"size:255; not null; unique"`
	Url       string    `gorm:"size:255; not null; unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
