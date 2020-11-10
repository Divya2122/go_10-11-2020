package main  
import (  
   "fmt"  
)  
type person struct {  
   firstName string  
   lastName  string  
   age       int  
}  
func main() {  
   a := person{age: 21, firstName: "Divya", lastName: "Avula", }  
   fmt.Println(a)  
   fmt.Println(a.firstName)  
}  