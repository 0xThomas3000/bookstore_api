package bookentity

import (
	"errors"
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
	Language      string      `json:"language" gorm:"language;"`
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
	b.Isbn = strings.TrimSpace(b.Isbn)

	if b.Title == "" {
		return ErrTitleIsEmpty
	}
	if b.Author == "" {
		return ErrAuthorIsEmpty
	}
	if b.Isbn == "" {
		return ErrIsbnIsEmpty
	}

	return nil
}

type BookUpdate struct {
	Title         *string     `json:"title" gorm:"column:title;"`
	Author        *string     `json:"author" gorm:"column:author;"`
	PublishedDate string      `json:"published_date" gorm:"column:published_date;"`
	Isbn          *string     `json:"isbn" gorm:"column:isbn;"`
	NumberOfPages *int        `json:"number_of_pages" gorm:"column:number_of_pages;"`
	CoverImage    *core.Image `json:"cover_image" gorm:"cover_image;"`
	Language      *string     `json:"language" gorm:"language;"`
}

func (BookUpdate) TableName() string {
	return Book{}.TableName()
}

var (
	ErrTitleIsEmpty  = errors.New("title cannot be empty")
	ErrAuthorIsEmpty = errors.New("author cannot be empty")
	ErrIsbnIsEmpty   = errors.New("isbn cannot be empty")
)
