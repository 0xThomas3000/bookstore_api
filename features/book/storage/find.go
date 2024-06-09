package bookstorage

import (
	"context"

	"github.com/0xThomas3000/bookstore_api/core"
	bookEntity "github.com/0xThomas3000/bookstore_api/features/book/entities"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*bookEntity.Book, error) {
	var data bookEntity.Book

	// 1: Not found error -> return gorm.ErrRecordNotFound
	// 2: Another error (DB parse key wrong)
	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, core.BookNotFound
		}
		return nil, core.ErrDB(err)
	}
	return &data, nil
}
