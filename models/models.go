package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Model struct {
	ID         int        `gorm:"primary_key" json:"id"`
	CreatedOn  time.Time  `json:"-"`
	ModifiedOn time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}
