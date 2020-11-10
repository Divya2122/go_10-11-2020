

package main 

import "fmt"

 
func main() { 

	 
	mychnl := make(chan string) 

	
	go func() { 
		mychnl <- "DIV"
		mychnl <- "div"
		mychnl <- "divys"
		mychnl <- "Divyareddy"
		close(mychnl) 
	}() 

	 
	for res := range mychnl { 
		fmt.Println(res) 
	} 
} 
