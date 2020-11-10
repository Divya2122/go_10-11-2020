
package main 

import ( 
	"fmt"
	"time"
) 

 
func Aname() { 

	arr1 := [4]string{"Divya", "rohith", "chandu", "babu"} 

	for t1 := 0; t1 <= 3; t1++ { 
	
		time.Sleep(150 * time.Millisecond) 
		fmt.Printf("%s\n", arr1[t1]) 
	} 
} 

 
func Aid() { 

	arr2 := [4]int{105, 107, 108, 109} 
	
	for t2 := 0; t2 <= 3; t2++ { 
	
		time.Sleep(500 * time.Millisecond) 
		fmt.Printf("%d\n", arr2[t2]) 
	} 
} 


func main() { 

	fmt.Println("!...Main Go-routine Start...!") 
 
	go Aname() 

	
	go Aid() 

	time.Sleep(3500 * time.Millisecond) 
	fmt.Println("\n!...Main Go-routine End...!") 
} 
