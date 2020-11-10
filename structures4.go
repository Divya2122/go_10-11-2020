package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books    
   var Book2 Books 
   var Book3 Books
   var Book4 Books   
 
   
   Book1.title = "Go Programming"
   Book1.author = "Mahesh Kumar"
   Book1.subject = "Go Programming Tutorial"
   Book1.book_id = 6495407


   Book2.title = "java Programming"
   Book2.author = "james gosling"
   Book2.subject = "java Programming Tutorial"
   Book2.book_id = 098765
   


   Book3.title = "DBMS"
   Book3.author = "DIVYA"
   Book3.subject = "DBMS Programming Tutorial"
   Book3.book_id = 123456

   
   Book4.title = "Telecom Billing"
   Book4.author = "Zara Ali"
   Book4.subject = "Telecom Billing Tutorial"
   Book4.book_id = 645700
 
  
   printBook(Book1)

   
   printBook(Book2)
   printBook(Book3)
   printBook(Book4)
}
func printBook( book Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}