package repositories

import (
	"context"
	"database/sql"

	"github.com/Alfian57/belajar_golang_restfull_api/models/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId uint64) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
