package bookstorage

import (
	"context"

	"github.com/0xThomas3000/bookstore_api/core"
	bookentity "github.com/0xThomas3000/bookstore_api/features/book/entities"
)

func (s *sqlStore) Add(context context.Context, data *bookentity.BookAdd) error {
	if err := s.db.Create(&data).Error; err != nil {
		return core.ErrDB(err) // Wrap sensitive error at Storage layer
	}
	return nil
}
