package bookentity

import (
	"errors"
	"regexp"
	"strings"

	"github.com/0xThomas3000/bookstore_api/core"
)

const EntityName = "Book"

type Book struct {
	core.SQLModel `json:",inline"`
	Title         string      `json:"title" gorm:"column:title;"`
	Author        string      `json:"author" gorm:"column:author;"`
	PublishedDate string      `json:"published_date" gorm:"column:published_date;"`
	Isbn          string      `json:"isbn" gorm:"column:isbn;"`
	NumberOfPages int         `json:"number_of_pages" gorm:"column:number_of_pages;"`
	CoverImage    *core.Image `json:"cover_image" gorm:"cover_image;"`
	Language      string      `json:"language" gorm:"language;"`
}

func (Book) TableName() string {
	return "books"
}

func (b *Book) Mask() {
	b.GenUID(core.DbTypeBook)
}

type BookAdd struct {
	core.SQLModel `json:",inline"`
	Title         string      `json:"title" gorm:"column:title;"`
	Author        string      `json:"author" gorm:"column:author;"`
	PublishedDate string      `json:"published_date" gorm:"column:published_date;"`
	Isbn          string      `json:"isbn" gorm:"column:isbn;"`
	NumberOfPages int         `json:"number_of_pages" gorm:"column:number_of_pages;"`
	CoverImage    *core.Image `json:"cover_image" gorm:"cover_image;"`
	Language      string      `json:"language" gorm:"column:language;default:English;"`
}

func (BookAdd) TableName() string {
	return Book{}.TableName()
}

func (b *BookAdd) Mask() {
	b.GenUID(core.DbTypeBook)
}

func (b *BookAdd) Validate() error {
	b.Title = strings.TrimSpace(b.Title)
	b.Author = strings.TrimSpace(b.Author)
	b.PublishedDate = strings.TrimSpace(b.PublishedDate)
	b.Isbn = strings.TrimSpace(b.Isbn)
	b.Language = strings.TrimSpace(b.Language)

	if b.Title == "" {
		return ErrTitleIsEmpty
	}
	if b.Author == "" {
		return ErrAuthorIsEmpty
	}
	if ok, _ := regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}$", b.PublishedDate); !ok {
		return ErrPublishedDateInvalid
	}
	if ok, _ := regexp.MatchString("^\\d{10,13}$", b.Isbn); !ok {
		return ErrIsbnInvalid
	}
	if b.NumberOfPages <= 0 {
		return ErrNumberInvalid
	}
	if b.Language == "" {
		return ErrLanguageIsEmpty
	}

	return nil
}

type BookUpdate struct {
	Title         *string     `json:"title" gorm:"column:title;"`
	Author        *string     `json:"author" gorm:"column:author;"`
	PublishedDate *string     `json:"published_date" gorm:"column:published_date;"`
	Isbn          *string     `json:"isbn" gorm:"column:isbn;"`
	NumberOfPages *int        `json:"number_of_pages" gorm:"column:number_of_pages;"`
	CoverImage    *core.Image `json:"cover_image" gorm:"cover_image;"`
	Language      *string     `json:"language" gorm:"language;"`
}

func (BookUpdate) TableName() string {
	return Book{}.TableName()
}

func (b *BookUpdate) Validate() error {
	if b.Title != nil {
		*b.Title = strings.TrimSpace(*b.Title)
	}
	if b.Author != nil {
		*b.Author = strings.TrimSpace(*b.Author)
	}
	if b.PublishedDate != nil {
		*b.PublishedDate = strings.TrimSpace(*b.PublishedDate)
	}
	if b.Isbn != nil {
		*b.Isbn = strings.TrimSpace(*b.Isbn)
	}
	if b.Language != nil {
		*b.Language = strings.TrimSpace(*b.Language)
	}
	if *b.Title == "" {
		return ErrTitleIsEmpty
	}
	if *b.Author == "" {
		return ErrAuthorIsEmpty
	}
	if ok, _ := regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}$", *b.PublishedDate); !ok {
		return ErrPublishedDateInvalid
	}
	if ok, _ := regexp.MatchString("^\\d{10,13}$", *b.Isbn); !ok {
		return ErrIsbnInvalid
	}
	if *b.NumberOfPages <= 0 {
		return ErrNumberInvalid
	}
	if *b.Language == "" {
		return ErrLanguageIsEmpty
	}

	return nil
}

var (
	ErrTitleIsEmpty         = errors.New("title cannot be empty")
	ErrAuthorIsEmpty        = errors.New("author cannot be empty")
	ErrPublishedDateInvalid = errors.New("published date must be in 'YYYY-MM-DD' format")
	ErrIsbnInvalid          = errors.New("isbn can only contain 10-13 digits")
	ErrNumberInvalid        = errors.New("number of pages must be greater than 0")
	ErrLanguageIsEmpty      = errors.New("language should not be empty")
)
