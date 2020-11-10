 
package main 

import "fmt"

 
type Employee struct { 
	firstName, lastName string 
	age, salary int
} 

 
func main() { 

	
	emp1 := &Employee{"Divya", "avula", 21, 25000}
        emp2 := &Employee{"Divya", "avula", 21, 25000}
        emp3 := &Employee{"Divya", "avula", 21, 25000} 

	 
	fmt.Println("First Name: ", emp1.firstName) 
	fmt.Println("Age: ", emp1.age) 
        fmt.Println("First Name: ", emp2.firstName) 
	fmt.Println("Age: ", emp2.age) 
        fmt.Println("First Name: ", emp3.firstName) 
	fmt.Println("Age: ", emp3.age) 
} 
