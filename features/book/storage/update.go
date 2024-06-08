package bookstorage

import (
	"context"

	"github.com/0xThomas3000/bookstore_api/core"
	bookEntity "github.com/0xThomas3000/bookstore_api/features/book/entities"
)

func (s *sqlStore) Update(
	context context.Context,
	id int,
	data *bookEntity.BookUpdate,
) error {
	if err := s.db.Table(bookEntity.BookUpdate{}.TableName()).
		Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return core.ErrDB(err)
	}

	return nil
}
