package bookbusiness

import (
	"context"
	"errors"
	"testing"

	"github.com/0xThomas3000/bookstore_api/core"
	bookentity "github.com/0xThomas3000/bookstore_api/features/book/entities"
)

type mockAddStore struct{}

func (mockAddStore) Add(context context.Context, data *bookentity.BookAdd) error {
	if data.Title == "Harry Potter" {
		return core.ErrDB(errors.New("invalid input data"))
	}
	data.Id = 10
	return nil
}

func TestNewAddBookBusiness(t *testing.T) {
	business := NewAddBookBusiness(mockAddStore{})

	// Case 1
	dataTest := bookentity.BookAdd{Title: ""}
	err := business.AddBook(context.Background(), &dataTest)
	if err == nil || err.Error() != "invalid input data" {
		t.Errorf("Failed")
	}

	// Case 2
	dataTest = bookentity.BookAdd{Title: "Harry Potter", Author: ""}
	err = business.AddBook(context.Background(), &dataTest)
	if err == nil || err.Error() != "invalid input data" {
		t.Errorf("Failed")
	}

	// Case 3
	dataTest = bookentity.BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997/06/26"}
	err = business.AddBook(context.Background(), &dataTest)
	if err == nil || err.Error() != "invalid input data" {
		t.Errorf("Failed")
	}

	// Case 4
	dataTest = bookentity.BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997-06-26",
		Isbn: "123456"}
	err = business.AddBook(context.Background(), &dataTest)
	if err == nil || err.Error() != "invalid input data" {
		t.Errorf("Failed")
	}

	// Case 5
	dataTest = bookentity.BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997-06-26",
		Isbn: "1234567890123", NumberOfPages: -3}
	err = business.AddBook(context.Background(), &dataTest)
	if err == nil || err.Error() != "invalid input data" {
		t.Errorf("Failed")
	}

	// Case 6
	dataTest = bookentity.BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997-06-26",
		Isbn: "1234567890123", NumberOfPages: 3, Language: ""}
	err = business.AddBook(context.Background(), &dataTest)
	if err == nil || err.Error() != "invalid input data" {
		t.Errorf("Failed")
	}

	// Case 7
	dataTest = bookentity.BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997-06-26",
		Isbn: "1234567890123", NumberOfPages: 3, Language: "English"}
	err = business.AddBook(context.Background(), &dataTest)
	if err == nil || err.Error() != "invalid input data" {
		t.Errorf("Failed")
	}
}
