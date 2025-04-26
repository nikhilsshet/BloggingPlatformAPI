package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Post struct {
    ID        uint             `gorm:"primaryKey" json:"id"`
    Title     string           `json:"title"`
    Content   string           `json:"content"`
    Category  string           `json:"category"`
    Tags      datatypes.JSON   `json:"tags"`
    CreatedAt time.Time        `json:"created_at"`
    UpdatedAt time.Time        `json:"updated_at"`
    DeletedAt gorm.DeletedAt   `gorm:"index" json:"-"`
}
