package main

import "fmt"

func main() {
    fmt.Println("Запуск системы управления библиотекой...")
    // 1. Создаем экземпляр библиотеки
    myLibrary := &Library{} //Пустая библиотека готова к работе
    fmt.Println("\n--- Наполняем библиотеку ---")

    //2. Добавляем читателей
    myLibrary.AddReader("Fatima", "Tokazova")
    myLibrary.AddReader("Sofia", "Tedeeva")

    //3.Добавляем книги
    myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
    myLibrary.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)

    fmt.Println("\n--- Библиотека готова к работе ---")
    fmt.Println("Количество читателей:", len(myLibrary.Readers))
    fmt.Println("Количество книг:", len(myLibrary.Books))

    fmt.Println("\nСписок всех книг в библиотеке:")
    myLibrary.ListAllBooks()
}



