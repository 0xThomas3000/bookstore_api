package bookbusiness

import (
	"context"
	"testing"

	bookentity "github.com/0xThomas3000/bookstore_api/features/book/entities"
)

type mockAddStore struct{}

func (mockAddStore) Add(context context.Context, data *bookentity.BookAdd) error {
	if data.Title == "" {
		return bookentity.ErrTitleIsEmpty
	}
	if data.Title == "Harry Potter" && data.Author == "" {
		return bookentity.ErrAuthorIsEmpty
	}
	if data.Title == "Harry Potter" && data.Author == "J. K. Rowling" && data.PublishedDate == "1997/06/26" {
		return bookentity.ErrPublishedDateInvalid
	}
	if data.Title == "Harry Potter" && data.Author == "J. K. Rowling" && data.PublishedDate == "1997-06-26" && data.Isbn == "123456" {
		return bookentity.ErrIsbnInvalid
	}
	if data.Title == "Harry Potter" && data.Author == "J. K. Rowling" &&
		data.PublishedDate == "1997-06-26" && data.Isbn == "1234567890123" && data.NumberOfPages == -3 {
		return bookentity.ErrNumberInvalid
	}
	if data.Title == "Harry Potter" && data.Author == "J. K. Rowling" &&
		data.PublishedDate == "1997-06-26" && data.Isbn == "1234567890123" &&
		data.NumberOfPages == 3 && data.Language == "" {
		return bookentity.ErrLanguageIsEmpty
	}
	if data.Title == "Harry Potter" && data.Author == "J. K. Rowling" &&
		data.PublishedDate == "1997-06-26" && data.Isbn == "1234567890123" &&
		data.NumberOfPages == 3 && data.Language == "English" {
		return nil
	}

	data.Id = 10
	return nil
}

func TestNewAddBookBusiness(t *testing.T) {
	business := NewAddBookBusiness(mockAddStore{})

	// Case 1
	dataTest := bookentity.BookAdd{Title: ""}
	err := business.AddBook(context.Background(), &dataTest)
	if err == nil || err.Error() != bookentity.ErrTitleIsEmpty.Error() {
		t.Errorf("Failed")
	}

	// Case 2
	dataTest = bookentity.BookAdd{Title: "Harry Potter", Author: ""}
	err = business.AddBook(context.Background(), &dataTest)
	if err == nil || err.Error() != bookentity.ErrAuthorIsEmpty.Error() {
		t.Errorf("Failed")
	}

	// Case 3
	dataTest = bookentity.BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997/06/26"}
	err = business.AddBook(context.Background(), &dataTest)
	if err == nil || err.Error() != bookentity.ErrPublishedDateInvalid.Error() {
		t.Errorf("Failed")
	}

	// Case 4
	dataTest = bookentity.BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997-06-26",
		Isbn: "123456"}
	err = business.AddBook(context.Background(), &dataTest)
	if err == nil || err.Error() != bookentity.ErrIsbnInvalid.Error() {
		t.Errorf("Failed")
	}

	// Case 5
	dataTest = bookentity.BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997-06-26",
		Isbn: "1234567890123", NumberOfPages: -3}
	err = business.AddBook(context.Background(), &dataTest)
	if err == nil || err.Error() != bookentity.ErrNumberInvalid.Error() {
		t.Errorf("Failed")
	}

	// Case 6
	dataTest = bookentity.BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997-06-26",
		Isbn: "1234567890123", NumberOfPages: 3, Language: ""}
	err = business.AddBook(context.Background(), &dataTest)
	if err == nil || err.Error() != bookentity.ErrLanguageIsEmpty.Error() {
		t.Errorf("Failed")
	}

	// Case 7
	dataTest = bookentity.BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997-06-26",
		Isbn: "1234567890123", NumberOfPages: 3, Language: "English"}
	err = business.AddBook(context.Background(), &dataTest)
	if err != nil {
		t.Errorf("Failed")
	}
}
