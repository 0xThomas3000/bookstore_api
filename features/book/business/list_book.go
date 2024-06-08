package bookbusiness

import (
	"context"

	"github.com/0xThomas3000/bookstore_api/core"
	bookEntity "github.com/0xThomas3000/bookstore_api/features/book/entities"
)

type ListBookStore interface {
	ListDataWithCondition(
		context context.Context,
		paging *core.Paging,
		moreKeys ...string,
	) ([]bookEntity.Book, error)
}

type listBookBiz struct {
	store ListBookStore
}

func NewListBookBusiness(store ListBookStore) *listBookBiz {
	return &listBookBiz{store: store}
}

func (biz *listBookBiz) ListBook(
	context context.Context,
	paging *core.Paging,
) ([]bookEntity.Book, error) {
	result, err := biz.store.ListDataWithCondition(context, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
