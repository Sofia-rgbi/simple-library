package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
    "simple-library/domain"
)

type Storable interface {
    Save() error
    Load() error
}

func SaveBookToCSV(filename string, books []domain.Book) error {
    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("не удалось создать файл %s: %w", filename, err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    headers := []string{"ID", "Название", "Автор", "Год"}
    if err := writer.Write(headers); err != nil {
        return fmt.Errorf("не удалось записать заголовок: %w", err)
    }

    for _, book := range books {
        record := []string{
            strconv.Itoa(book.ID),
            book.Title,
            book.Author,
            strconv.Itoa(book.Year),
        }

        if err := writer.Write(record); err != nil {
            return fmt.Errorf("не удалось записать книгу с ID %d: %w", book.ID, err)
        }
    }

    return nil
}

func LoadBookFromCSV(filename string) ([]domain.Book, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("не удалось открыть файл %s: %w", filename, err)
    }
    defer file.Close()

    reader := csv.NewReader(file)

    // Проверка заголовков
    _, err = reader.Read()
    if err != nil {
        return nil, fmt.Errorf("не удалось прочитать заголовок из файла: %w", err)
    }

    // Опционально: проверить, что заголовки корректны

    records, err := reader.ReadAll()
    if err != nil {
        return nil, fmt.Errorf("не удалось прочитать записи из файла: %w", err)
    }

    var books []domain.Book
    for _, record := range records {
        // Проверим, что в записи 4 поля
        if len(record) < 4 {
            return nil, fmt.Errorf("неверный формат записи: %v", record)
        }

        id, err := strconv.Atoi(record[0])
        if err != nil {
            return nil, fmt.Errorf("не удалось конвертировать ID: %w", err)
        }

        year, err := strconv.Atoi(record[3])
        if err != nil {
            return nil, fmt.Errorf("не удалось конвертировать год: %w", err)
        }

        book := domain.Book{
            ID:     id,
            Title:  record[1],
            Author: record[2],
            Year:   year,
        }

        books = append(books, book)
    }

    return books, nil
}
	 