

package main 

import "fmt"

 
type Address struct { 
	Name string 
	city string 
	Pincode int
} 

func main() { 

	var a Address 
	fmt.Println(a) 

	
	a1 := Address{"divya", "galiveedu", 3623572} 

	fmt.Println("Address1: ", a1) 

	
	a2 := Address{Name: "chandu", city: "madanapalle", 
								Pincode: 277001} 

	fmt.Println("Address2: ", a2) 

	
	a3 := Address{Name: "kholi"} 
	fmt.Println("Address3: ", a3) 
} 
