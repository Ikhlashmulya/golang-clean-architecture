package repository

import (
	"context"
	"database/sql"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/domain"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/exception"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}

func (repository *CategoryRepositoryImpl) Insert(ctx context.Context, category domain.Category) (lastInsertId int) {
	tx, err := repository.DB.Begin()
	exception.PanicIfError(err)
	defer commitOrRollback(tx)

	SQL := "INSERT INTO category (name) VALUES (?)"

	result, err := tx.ExecContext(ctx, SQL, category.Name)
	exception.PanicIfError(err)

	id, err := result.LastInsertId()
	exception.PanicIfError(err)

	lastInsertId = int(id)
	return
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, category domain.Category) {
	tx, err := repository.DB.Begin()
	exception.PanicIfError(err)
	defer commitOrRollback(tx)

	SQL := "UPDATE category SET name = ? WHERE id = ?"

	_, err = tx.ExecContext(ctx, SQL, category.Name, category.Id)
	exception.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := repository.DB.Begin()
	exception.PanicIfError(err)
	defer commitOrRollback(tx)

	SQL := "DELETE FROM category WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, categoryId)
	exception.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId int) (category domain.Category, err error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := repository.DB.QueryContext(ctx, SQL, categoryId)
	exception.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		errScan := rows.Scan(&category.Id, &category.Name)
		exception.PanicIfError(errScan)
		return category, nil
	} else {
		category = domain.Category{}
		err = sql.ErrNoRows
		return
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context) (responses []domain.Category) {
	SQL := "SELECT id, name FROM category"
	rows, err := repository.DB.QueryContext(ctx, SQL)
	exception.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		category := domain.Category{}
		errScan := rows.Scan(&category.Id, &category.Name)
		exception.PanicIfError(errScan)

		responses = append(responses, category)
	}

	return responses
}

func commitOrRollback(tx *sql.Tx) {
	err := recover()
	switch err {
	case nil:
		errCommit := tx.Commit()
		exception.PanicIfError(errCommit)
	default:
		errRollback := tx.Rollback()
		exception.PanicIfError(errRollback)
		panic(err)
	}
}
