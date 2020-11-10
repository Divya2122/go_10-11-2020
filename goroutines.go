
package main 

import "fmt"

func display(str string) { 
	for w := 0; w < 10; w++ { 
		fmt.Println(str) 
	} 
} 

func main() { 

	 
	go display("Welcome") 

	
	display("divya") 
} 
