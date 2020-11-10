


package main 

import "fmt"
a 

func main() { 

	
	mychnl := make(chan string, 4) 
	mychnl <- "Div"
	mychnl <- "div"
	mychnl <- "Divs"
	mychnl <- "divyareddy"

	
	fmt.Println("Length of the channel is: ", len(mychnl)) 
} 
