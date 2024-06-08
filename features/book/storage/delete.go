package bookstorage

import (
	"context"
	"github.com/0xThomas3000/bookstore_api/core"
	bookEntity "github.com/0xThomas3000/bookstore_api/features/book/entities"
)

func (s *sqlStore) Delete(
	context context.Context,
	id int,
	data *bookEntity.Book,
) error {
	if err := s.db.Table(bookEntity.Book{}.TableName()).
		Where("id = ?", id).
		Delete(&data).Error; err != nil {
		return core.ErrDB(err)
	}
	return nil
}
