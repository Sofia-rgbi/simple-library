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
	fmt.Printf("Читатель %s %s взял книгу '%s' (%s, %d)\n", r.FirstName, r.LastName, b.Title, b.Author)
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
