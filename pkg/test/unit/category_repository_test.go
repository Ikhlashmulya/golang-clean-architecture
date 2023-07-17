package test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/domain"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestInsertRepository(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		categoryRepository := repository.NewCategoryRepository(db)

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO category").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		category := domain.Category{
			Name: "category1",
		}

		gotLastInsertId := categoryRepository.Insert(context.Background(), category)

		assert.Equal(t, 1, gotLastInsertId)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failed", func(t *testing.T) {
		assert.Panics(t, func() {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			categoryRepository := repository.NewCategoryRepository(db)

			mock.ExpectBegin()
			mock.ExpectExec("INSERT INTO category").WillReturnError(sql.ErrConnDone)
			mock.ExpectRollback()

			category := domain.Category{
				Name: "category1",
			}

			_ = categoryRepository.Insert(context.Background(), category)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	})
}

func TestUpdateRepository(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		categoryRepository := repository.NewCategoryRepository(db)

		category := domain.Category{
			Id:   1,
			Name: "category1",
		}

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE category").
			WithArgs(category.Name, category.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		categoryRepository.Update(context.Background(), category)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failed", func(t *testing.T) {
		assert.Panics(t, func() {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			categoryRepository := repository.NewCategoryRepository(db)

			category := domain.Category{
				Id:   1,
				Name: "category1",
			}

			mock.ExpectBegin()
			mock.ExpectExec("UPDATE category").
				WithArgs(category.Name, category.Id).
				WillReturnError(sql.ErrConnDone)
			mock.ExpectCommit()

			categoryRepository.Update(context.Background(), category)

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	})
}

func TestDeleteRepository(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		categoryRepository := repository.NewCategoryRepository(db)

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM category").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		categoryRepository.Delete(context.Background(), 1)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failed", func(t *testing.T) {
		assert.Panics(t, func() {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			categoryRepository := repository.NewCategoryRepository(db)

			mock.ExpectBegin()
			mock.ExpectExec("DELETE FROM category").WillReturnError(sql.ErrConnDone)
			mock.ExpectCommit()

			categoryRepository.Delete(context.Background(), 1)

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	})
}

func TestFindByIdRepository(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		categoryRepository := repository.NewCategoryRepository(db)

		row := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "category1")

		mock.ExpectQuery("SELECT id, name FROM category").
			WithArgs(1).
			WillReturnRows(row)

		gotCategory, err := categoryRepository.FindById(context.Background(), 1)
		assert.NoError(t, err)

		assert.Equal(t, 1, gotCategory.Id)
		assert.Equal(t, "category1", gotCategory.Name)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		categoryRepository := repository.NewCategoryRepository(db)

		mock.ExpectQuery("SELECT id, name FROM category").
			WithArgs(2).
			WillReturnRows(&sqlmock.Rows{})

		category, err := categoryRepository.FindById(context.Background(), 2)
		if assert.Error(t, err) {
			assert.Equal(t, "category is not found", err.Error())
		}
		assert.Empty(t, category)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestFindAllRepository(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		categoryRepository := repository.NewCategoryRepository(db)

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "category1").AddRow(2, "category2")

		mock.ExpectQuery("SELECT id, name FROM category").
			WillReturnRows(rows)

		gotCategories := categoryRepository.FindAll(context.Background())

		assert.Equal(t, 1, gotCategories[0].Id)
		assert.Equal(t, "category1", gotCategories[0].Name)
		assert.Equal(t, 2, gotCategories[1].Id)
		assert.Equal(t, "category2", gotCategories[1].Name)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("no content", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		categoryRepository := repository.NewCategoryRepository(db)

		mock.ExpectQuery("SELECT id, name FROM category").
			WillReturnRows(&sqlmock.Rows{})

		gotCategories := categoryRepository.FindAll(context.Background())
		assert.Empty(t, gotCategories)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
