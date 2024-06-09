package bookbusiness

import (
	"context"

	"github.com/0xThomas3000/bookstore_api/core"
	bookEntity "github.com/0xThomas3000/bookstore_api/features/book/entities"
)

type UpdateBookStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*bookEntity.Book, error)

	Update(
		context context.Context,
		id int,
		data *bookEntity.BookUpdate,
	) error
}

type updateBookBusiness struct {
	store UpdateBookStore
}

func NewUpdateBookBusiness(store UpdateBookStore) *updateBookBusiness {
	return &updateBookBusiness{store: store}
}

func (business *updateBookBusiness) UpdateBook(
	context context.Context,
	id int,
	data *bookEntity.BookUpdate,
) error {
	_, err := business.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return core.ErrEntityNotFound(bookEntity.EntityName, err)
	}

	if err := business.store.Update(context, id, data); err != nil {
		return core.ErrBadRequest(err)
	}

	return nil
}
