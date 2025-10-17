package main

import (
	"fmt"
)

func main() {
   

    user1 := Reader{
        ID:        1,
        FirstName: "Fatima",
        LastName:  "Tokazova",
        IsActive:  true,
    }
    user2 := Reader{
        ID:        2,
        FirstName: "Fatima",
        LastName:  "Tsikalova",
        IsActive:  true,
    }

    book1 := Book{
        Title:  "Idiot",
        Author: "Dostoevskiy",
        Year:   1869,
    }

    // Выдаем книгу конкретному читателю
    book1.IssuesBook(&user1)
    fmt.Println(book1.String()) // виден ID читателя

    // Пробуем выдать уже выданную книгу другому читателю
    book1.IssuesBook(&user2)

    // Возвращаем книгу
    book1.ReturnBook()
    fmt.Println(book1.String())

    // Выдаем книгу неактивному читателю
    user1.IsActive = false
    book1.IssuesBook(&user1)

    // Демонстрация AssignBook
    user2.AssignBook(&book1)

    notifiers := []Notifier{
        EmailNotifier{EmailAddress: "student@example.com"},
        SMSNotifier{PhoneNumber: "+79991234567"},
    }

    message := "Ваша книга просрочена!"

    for _, notifier := range notifiers {
        notifier.Notify(message)
    }
}
