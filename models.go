package main

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
	Issues bool
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
func (r* Reader) Deactivate() {
	r.IsActive = false
}

func (r Reader) String() string{
	status := ""

	if r.IsActive{
		status = "Активен"
	} else {
		status = "Не активен"
	}
	return fmt.Sprintf("Пользователь %s %s, ID: %d, пользователь:%", r.FirstName, r.LastName, r.ID, status)
}