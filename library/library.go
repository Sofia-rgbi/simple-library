package library

import (
	"fmt"
	"simple-library/domain"
)

type Library struct {
	Books        []*domain.Book
	Readers      []*domain.Reader
	lastBookID   int
	LastReaderID int
}

func (lib *Library) FindReaderById(id int) (any, error) {
	panic("unimplemented")
}

func (lib *Library) FindBookById(id int) (any, error) {
	panic("unimplemented")
}

func (lib *Library) AddReader(firstName, lastName string) *domain.Reader {
	lib.LastReaderID++

	newReader := &domain.Reader{
		ID:        lib.LastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}

	lib.Readers = append(lib.Readers, newReader)
	return newReader
}

func (lib *Library) AddBook(title, author string, year int) (*domain.Book, error) {

	for _, b := range lib.Books {
		if b.Title == title && b.Author == author {
			return nil, fmt.Errorf("книга '%s' автора '%s' уже существует", title, author)
		}
	}

	lib.lastBookID++

	newBook := &domain.Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false,
	}

	lib.Books = append(lib.Books, newBook)
	return newBook, nil
}

func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (lib *Library) FindReaderByID(id int) (*domain.Reader, error) {
	for _, reader := range lib.Readers {
		if reader.ID == id {
			return reader, nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

func (lib *Library) IssueBookToReader(bookID int, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}

	if book.IsIssued {
		return fmt.Errorf("книга %s уже выдана", book.Title)
	}

	if !reader.IsActive {
		return fmt.Errorf("читатель %s %s не активен", reader.FirstName, reader.LastName)
	}

	book.IsIssued = true
	book.ReaderId = &reader.ID

	return nil
}

func (lib *Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	return book.ReturnBook()
}

func New() *Library {
	return &Library{
		Books:   []*domain.Book{},
		Readers: []*domain.Reader{},
	}
}
