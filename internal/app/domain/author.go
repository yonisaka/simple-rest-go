package domain

import (
	"context"
	"time"
)

type Author struct {
	ID        int64     `json="id"`
	Name      string    `json="name" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `json="created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json="updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (Author, error)
}
