package main

   import (
       "fmt"
       "time"
   )

   
   func printNumbers(id int, n int) {
	      for i := 1; i <= n; i++ {
            fmt.Prinf("Worker %d number %d\n", id, i)
            time.Sleep(100 * time.Millisecond)
          }
   }

   func main() {
	   fmt.Println("Start lab1")
       printNumbers(1, 5)
       fmt.Println("End lab1")

   }

