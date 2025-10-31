package main

import (
    "flag"
    "fmt"
    "os"
    "simple-library/library"
)

func main() {
    saveCmd := flag.NewFlagSet("save", flag.ExitOnError)
    saveFile := saveCmd.String("file", "books.csv", "файл для сохранения")

    loadCmd := flag.NewFlagSet("load", flag.ExitOnError)
    loadFile := loadCmd.String("file", "books.csv", "файл для загрузки")

    if len(os.Args) < 2 {
        fmt.Println("Ожидалась команда: save или load")
        os.Exit(1)
    }

    myLibrary := library.New()
    myLibrary.AddBook("Преступление и наказание", "Федор Достоевский", 1866)
    myLibrary.AddReader("Fatima", "Tokazova")
    myLibrary.IssueBookToReader(1, 1)

    switch os.Args[1] {
    case "save":
        saveCmd.Parse(os.Args[2:])
        // здесь вызываем функцию сохранения, например myLibrary.SaveToCSV(*saveFile)
        fmt.Printf("Сохраняем библиотеку в файл %s\n", *saveFile)
        // Добавь код сохранения

    case "load":
        loadCmd.Parse(os.Args[2:])
        // здесь вызываем функцию загрузки, например myLibrary.LoadFromCSV(*loadFile)
        fmt.Printf("Загружаем библиотеку из файла %s\n", *loadFile)
        // Добавь код загрузки

    default:
        fmt.Println("Неизвестная команда. Используйте save или load.")
        os.Exit(1)
    }

    if len(myLibrary.Books) > 0 && len(myLibrary.Readers) > 0 {
        fmt.Printf("Книга %s, автор %s, год издания %d\n",
            myLibrary.Books[0].Title,
            myLibrary.Books[0].Author,
            myLibrary.Books[0].Year)
        fmt.Printf("Читатель: %s %s\n",
            myLibrary.Readers[0].FirstName,
            myLibrary.Readers[0].LastName)
    }
}