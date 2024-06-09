package bookbusiness

import (
	"context"

	"github.com/0xThomas3000/bookstore_api/core"
	bookentity "github.com/0xThomas3000/bookstore_api/features/book/entities"
)

type AddBookStore interface {
	Add(context context.Context, data *bookentity.BookAdd) error
}

type addBookBusiness struct {
	store AddBookStore
}

func NewAddBookBusiness(store AddBookStore) *addBookBusiness {
	return &addBookBusiness{store: store}
}

func (business *addBookBusiness) AddBook(context context.Context, data *bookentity.BookAdd) error {
	if err := data.Validate(); err != nil {
		return core.ErrBadRequest(err)
	}

	if err := business.store.Add(context, data); err != nil {
		return core.ErrCannotCreateEntity(bookentity.EntityName, err)
	}

	return nil
}
