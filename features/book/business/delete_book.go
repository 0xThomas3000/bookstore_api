package bookbusiness

import (
	"context"
	"github.com/0xThomas3000/bookstore_api/core"
	bookEntity "github.com/0xThomas3000/bookstore_api/features/book/entities"
)

type DeleteBookStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*bookEntity.Book, error)

	Delete(
		context context.Context,
		id int,
		data *bookEntity.Book,
	) error
}

type deleteBookBusiness struct {
	store DeleteBookStore
}

func NewDeleteBookBusiness(store DeleteBookStore) *deleteBookBusiness {
	return &deleteBookBusiness{store: store}
}

func (business *deleteBookBusiness) DeleteBook(
	context context.Context,
	id int,
	data *bookEntity.Book,
) error {
	_, err := business.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return core.ErrEntityNotFound(bookEntity.EntityName, err)
	}

	if err := business.store.Delete(context, id, data); err != nil {
		return core.ErrCannotDeleteEntity(bookEntity.EntityName, nil)
	}

	return nil
}
