package health

import (
	"context"

	"gorm.io/gorm"
)

type RepositoryInterface interface {
	CheckUpTimeDB(ct context.Context) (err error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) CheckUpTimeDB(ctx context.Context) (err error) {
	db, err := r.db.WithContext(ctx).DB()
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return
	}

	return
}
