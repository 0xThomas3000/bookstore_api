package bookentity

import (
	"testing"
)

type TestBookAdd struct {
	Input       BookAdd
	ErrExpected error
}

type TestBookUpdate struct {
	Input       BookUpdate
	ErrExpected error
}

func TestBookAdd_Validate(t *testing.T) {
	dataTable := []TestBookAdd{
		{Input: BookAdd{Title: ""}, ErrExpected: ErrTitleIsEmpty},
		{Input: BookAdd{Title: "Harry Potter", Author: ""}, ErrExpected: ErrAuthorIsEmpty},
		{Input: BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997/06/26"},
			ErrExpected: ErrPublishedDateInvalid},
		{Input: BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997-06-26",
			Isbn: "123456"}, ErrExpected: ErrIsbnInvalid},
		{Input: BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997-06-26",
			Isbn: "1234567890123", NumberOfPages: -3}, ErrExpected: ErrNumberInvalid},
		{Input: BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997-06-26",
			Isbn: "1234567890123", NumberOfPages: 3, Language: ""}, ErrExpected: ErrLanguageIsEmpty},
		{Input: BookAdd{Title: "Harry Potter", Author: "J. K. Rowling", PublishedDate: "1997-06-26",
			Isbn: "1234567890123", NumberOfPages: 3, Language: "English"}, ErrExpected: nil},
	}

	for _, item := range dataTable {
		err := item.Input.Validate()
		if err != item.ErrExpected {
			t.Errorf("Book validation status; Input: %v, Expect: %v, Output: %v", item.Input, item.ErrExpected, err)
		}
	}
}

func TestBookUpdate_Validate(t *testing.T) {
	var title1, title2 = "", "Head First Java"
	var author1, author2 = "", "Kathy Sierra"
	var pubDate1, pubDate2 = "2005/05/15", "2005-05-15"
	var isbn1, isbn2 = "9999", "1234567890123"
	var number1, number2 = -3, 3
	var language1, language2 = "", "English"

	dataTable := []TestBookUpdate{
		{Input: BookUpdate{Title: &title1}, ErrExpected: ErrTitleIsEmpty},
		{Input: BookUpdate{Title: &title2, Author: &author1}, ErrExpected: ErrAuthorIsEmpty},
		{Input: BookUpdate{Title: &title2, Author: &author2, PublishedDate: &pubDate1},
			ErrExpected: ErrPublishedDateInvalid},
		{Input: BookUpdate{Title: &title2, Author: &author2, PublishedDate: &pubDate2,
			Isbn: &isbn1}, ErrExpected: ErrIsbnInvalid},
		{Input: BookUpdate{Title: &title2, Author: &author2, PublishedDate: &pubDate2,
			Isbn: &isbn2, NumberOfPages: &number1}, ErrExpected: ErrNumberInvalid},
		{Input: BookUpdate{Title: &title2, Author: &author2, PublishedDate: &pubDate2,
			Isbn: &isbn2, NumberOfPages: &number2, Language: &language1}, ErrExpected: ErrLanguageIsEmpty},
		{Input: BookUpdate{Title: &title2, Author: &author2, PublishedDate: &pubDate2,
			Isbn: &isbn2, NumberOfPages: &number2, Language: &language2}, ErrExpected: nil},
	}

	for _, item := range dataTable {
		err := item.Input.Validate()
		if err != item.ErrExpected {
			t.Errorf("Book validation status; Input: %v, Expect: %v, Output: %v", item.Input, item.ErrExpected, err)
		}
	}
}
