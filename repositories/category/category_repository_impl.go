package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Alfian57/belajar_golang_restfull_api/helper"
	"github.com/Alfian57/belajar_golang_restfull_api/models/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = uint64(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId uint64) (domain.Category, error) {
	SQL := "SELECT id, name FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		rows.Scan(&category.Id, &category.Name)
		return category, nil
	}

	return category, errors.New("category not found")
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		rows.Scan(&category.Id, &category.Name)

		categories = append(categories, category)
	}

	return categories
}
