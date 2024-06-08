package bookstorage

import (
	"context"
	"github.com/0xThomas3000/bookstore_api/core"
	bookEntity "github.com/0xThomas3000/bookstore_api/features/book/entities"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	paging *core.Paging,
	moreKeys ...string,
) ([]bookEntity.Book, error) {
	var result []bookEntity.Book

	db := s.db.Table(bookEntity.Book{}.TableName())

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, core.ErrDB(err)
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := core.FromBase58(v)
		if err != nil {
			return nil, core.ErrDB(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, core.ErrDB(err)
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask()
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
