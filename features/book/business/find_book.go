package bookbusiness

import (
	"context"
	bookEntity "github.com/0xThomas3000/bookstore_api/features/book/entities"
)

type FindBookStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*bookEntity.Book, error)
}

type findBookBusiness struct {
	store FindBookStore
}

func NewFindBookBusiness(store FindBookStore) *findBookBusiness {
	return &findBookBusiness{store: store}
}

func (business *findBookBusiness) FindBook(context context.Context, id int) (*bookEntity.Book, error) {
	result, err := business.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}
	return result, nil
}
