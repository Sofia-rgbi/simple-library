package main

import (
    "fmt"
    "simple-library/library"


)

func main() {
   myLibrary := library.New()

   myLibrary.AddBook("Преступление и наказание","Федор Достоевский", 1886)
   myLibrary.AddReader("Fatima","Tokazova")
   myLibrary.IssueBookToReader(1,1)

   fmt.Printf("Книга %s, авторства %s, %d года находится у %s %s", myLibrary.Books[0].Title,myLibrary.Books[0].Author,myLibrary.Books[0].Year, myLibrary.Readers[0].FirstName,myLibrary.Readers[0].LastName,)

}