package main

import "fmt"

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderId *int
}


type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

// Library - наша центральная структура, агрегатор.
type Library struct {
    Books   []*Book
    Readers []*Reader

    // Счетчики для генерации уникальных ID
    lastBookID   int
    lastReaderID int
}


func (lib *Library) FindBookByID(id int) (*Book, error) {
    for _, book := range lib.Books {
        if book.ID == id {
            return book, nil
        }
    }
    return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (lib *Library) FindReaderByID(id int) (*Reader, error) {
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

    fmt.Printf("Книга %s была выдана читателю %s %s\n", book.Title, reader.FirstName, reader.LastName)
    return nil
}

func (lib *Library) ListAllBooks() {
	for _, book := range lib.Books{
		fmt.Println(book.String())
	}
	
	
}


func (lib *Library) AddReader(firstName, lastName string) *Reader {
    // Увеличиваем счетчик ID
    lib.lastReaderID++
    
    // Создаем нового читателя
    newReader := &Reader{
        ID:        lib.lastReaderID,
        FirstName: firstName,
	 LastName:  lastName,
        IsActive:  true, // Новый читатель всегда активен
    }

    // Добавляем читателя в срез
    lib.Readers = append(lib.Readers, newReader)

    fmt.Printf("Зарегистрирован новый читатель: %s\n", newReader)
    return newReader
}

// AddBook создает новую книгу и добавляет ее в библиотеку.
// Автоматически присваивает уникальный ID.
func (lib *Library) AddBook(title, author string, year int) *Book {
    lib.lastBookID++

    newBook := &Book{
        ID:       lib.lastBookID,
        Title:    title,
        Author:   author,
 Year:     year,
        IsIssued: false, // Новая книга всегда в наличии
    }

    lib.Books = append(lib.Books, newBook)

    fmt.Printf("Добавлена новая книга: %s\n", newBook)
    return newBook
}



//Выводит в консоль информацю о читателе
func (r Reader) DisplayReader() {
	fmt.Printf("Читатель: %s %s (ID: %d)\n", r.FirstName, r.LastName, r.ID)
}

//Deactivate делает читателя не активным
func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (r Reader) String() string {
	status := ""

	if r.IsActive {
		status = "Активен"
	} else {
		status = "Не активен"
	}
	return fmt.Sprintf("Пользователь %s %s, ID: %d, пользователь:%s", r.FirstName, r.LastName, r.ID, status)
}

func (b Book) String() string {

	status := "В библиотеке"
	if b.IsIssued && b.ReaderId != nil {
		status = fmt.Sprintf("На руках у читателя с ID: %d", *b.ReaderId)
	}
	return fmt.Sprintf(`"%s (%s, %d),Статус: %s "`, b.Title, b.Author, b.Year, status)
}

func (r *Reader) AssignBook(b *Book) {
	fmt.Printf("Читатель %s %s взял книгу '%s' (%s, %d)\n", r.FirstName, r.LastName, b.Title, b.Author, b.Year)
}

//IssueBook выдает книгу читателю
func (b *Book) IssuesBook(r *Reader) {
	if b.IsIssued {
		fmt.Printf("Книга %s уже кому-то выдана\n", b.Title)
		return
	}
	if !r.IsActive {
		fmt.Printf("Читатель %s %s не активен и не может получить книгу.", r.FirstName, r.LastName)
		return
	}

	b.IsIssued = true
	b.ReaderId = &r.ID
	fmt.Printf("Книга %s была выдана\n", b.Title)
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s и так в библиотеке", b.Title)
		return
	}
	b.IsIssued = false
	b.ReaderId = nil
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}
