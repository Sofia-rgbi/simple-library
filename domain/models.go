package domain

import "fmt"

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderId *int
}



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
	return fmt.Sprintf(`"%s (%s, %d), Статус: %s"`, b.Title, b.Author, b.Year, status)
}

func (r *Reader) AssignBook(b *Book) error {

	if b.IsIssued {
		return fmt.Errorf("книга '%s' уже выдана", b.Title)
	}
	if !r.IsActive {
		return fmt.Errorf("читатель %s %s не активен", r.FirstName, r.LastName)
	}
	
	b.IsIssued = true
	b.ReaderId = &r.ID
	return nil
}

func (b *Book) IssuesBook(r *Reader) error {

	if b.IsIssued {
		return fmt.Errorf("книга %s уже кому-то выдана", b.Title)
	}
	if !r.IsActive {
		return fmt.Errorf("читатель %s %s не активен и не может получить книгу", r.FirstName, r.LastName)
	}

	b.IsIssued = true
	b.ReaderId = &r.ID
	return nil
}

func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}

	b.IsIssued = false
	b.ReaderId = nil
	return nil
}

