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
    

    
fmt.Println("---Тестируем выдачу книг---")
//Выдаем книгу 1 читателю 1
err := myLibrary.IssueBookToReader(1, 1)
if err != nil {
	fmt.Println("Ошибка выдачи", err)
}

//Проверить статус книги после выдачи
book, _ := myLibrary.FindBookByID(1)
if book != nil {
	fmt.Println("Статус книги после выдачи:", book)
}

//Попытка выдать несуществующую книгу
err = myLibrary.IssueBookToReader(99, 1)
if err != nil {
	fmt.Println("Ожидаемая ошибка:", err)
}

err1 := myLibrary.ReturnBook(1)
if err1 != nil {
    fmt.Println("Ошибка при выдаче книги: ", err)
} else {
    fmt.Println("Книга возвращена успешно")
}

err2 := myLibrary.ReturnBook(1)
if err2 != nil {
    fmt.Println("Ошибка при выдаче книги: ", err)
} else {
    fmt.Println("Книга возвращена успешно")
}

config := map[string]string {
    "PORT" : "2123",
}


//config1 := map[string]string {
//    "NOPORT" : "2132",
//}
//
port, err := GetPortFromConfig(config)
if err != nil {
    fmt.Println("Произошла ошибка чтения порта:", err)
}else {
    fmt.Println("Порт номер:", port)
}



}