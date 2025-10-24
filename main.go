package main

import "fmt"

func main() {
    fmt.Println("Запуск системы управления библиотекой...")

    // Создаем экземпляр библиотеки
    myLibrary := &Library{}

    fmt.Println("\n--- Наполняем библиотеку ---")

    // Добавляем читателей
    myLibrary.AddReader("Fatima", "Tokazova")
    myLibrary.AddReader("Sofia", "Tedeeva")

    // Добавляем книги
    myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
    myLibrary.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)

    fmt.Println("\n--- Библиотека готова к работе ---")
    fmt.Println("Количество читателей:", len(myLibrary.Readers))
    fmt.Println("Количество книг:", len(myLibrary.Books))

    fmt.Println("\n--- Список всех книг в библиотеке ---")

    // Сценарий 1: Успешная выдача книги
    fmt.Println("\n--- Сценарий 1: Успешная выдача книги ---")
    err := myLibrary.IssueBookToReader(1, 1)
    if err != nil {
        fmt.Println("Ошибка выдачи книги:", err)
    } else {
        fmt.Println("Книга успешно выдана читателю")
        
        // Проверяем статус книги после выдачи
        book, _ := myLibrary.FindBookByID(1)
        if book != nil {
            fmt.Println("Статус книги после выдачи:", book)
        }
    }

    // Сценарий 2: Попытка выдать уже выданную книгу
    fmt.Println("\n--- Сценарий 2: Попытка выдать уже выданную книгу ---")
    err = myLibrary.IssueBookToReader(1, 2)
    if err != nil {
        fmt.Println("Ошибка выдачи книги:", err)
    } else {
        fmt.Println("Книга успешно выдана читателю")
    }

    // Сценарий 3: Попытка выдать книгу несуществующему читателю
    fmt.Println("\n--- Сценарий 3: Попытка выдать книгу несуществующему читателю ---")
    err = myLibrary.IssueBookToReader(1, 99)
    if err != nil {
        fmt.Println("Ошибка выдачи книги:", err)
    } else {
        fmt.Println("Книга успешно выдана читателю")
    }

    // Сценарий 4: Попытка выдать несуществующую книгу
    fmt.Println("\n--- Сценарий 4: Попытка выдать несуществующую книгу ---")
    err = myLibrary.IssueBookToReader(99, 1)
    if err != nil {
        fmt.Println("Ошибка выдачи книги:", err)
    } else {
        fmt.Println("Книга успешно выдана читателю")
    }

    // Сценарий 5: Успешный возврат книги
    fmt.Println("\n--- Сценарий 5: Успешный возврат книги ---")
    err = myLibrary.ReturnBook(1)
    if err != nil {
        fmt.Println("Ошибка возврата книги:", err)
    } else {
        fmt.Println("Книга успешно возвращена в библиотеку")
        
        // Проверяем статус книги после возврата
        book, _ := myLibrary.FindBookByID(1)
        if book != nil {
            fmt.Println("Статус книги после возврата:", book)
        }
    }

    // Сценарий 6: Попытка вернуть книгу, которая уже в библиотеке
    fmt.Println("\n--- Сценарий 6: Попытка вернуть книгу, которая уже в библиотеке ---")
    err = myLibrary.ReturnBook(1)
    if err != nil {
        fmt.Println("Ошибка возврата книги:", err)
    } else {
        fmt.Println("Книга успешно возвращена в библиотеку")
    }

    // Сценарий 7: Попытка вернуть несуществующую книгу
    fmt.Println("\n--- Сценарий 7: Попытка вернуть несуществующую книгу ---")
    err = myLibrary.ReturnBook(99)
    if err != nil {
        fmt.Println("Ошибка возврата книги:", err)
    } else {
        fmt.Println("Книга успешно возвращена в библиотеку")
    }

    // Тестирование функции GetPortFromConfig
    fmt.Println("\n--- Тестирование функции GetPortFromConfig ---")
    
    // Тест 1: Конфиг без порта
    config1 := map[string]string{
        "NOPORT": "2132",
    }
    port, err := GetPortFromConfig(config1)
    if err != nil {
        fmt.Println("Ошибка чтения порта:", err)
    } else {
        fmt.Println("Порт номер:", port)
    }

    // Тест 2: Конфиг с портом
    config2 := map[string]string{
        "PORT": "8080",
    }
    port, err = GetPortFromConfig(config2)
    if err != nil {
        fmt.Println("Ошибка чтения порта:", err)
    } else {
        fmt.Println("Порт номер:", port)
    }

    fmt.Println("\n--- Работа системы управления библиотекой завершена ---")
}