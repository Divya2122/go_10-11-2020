package main 
  
import "fmt"
  
type Address struct { 
    Name    string 
    city    string 
    state   string
    Pincode int
} 
  
func main() { 
  
    
    var a Address  
    fmt.Println(a)
    
    var b Address  
    fmt.Println(b) 
  
    
    a1 := Address{"Divya", "madanapalle", "Ap", 3623572} 
  
    fmt.Println("Address1: ", a1) 
   

    b1 := Address{"gopi", "kadapa", "Ap", 123456} 
  
    fmt.Println("Address1: ", b1)
  
    
    a2 := Address{Name: "chandu", city: "galiveedu", state: "Ap", 
                                 Pincode: 277001} 
  
    fmt.Println("Address2: ", a2) 
  
    
    a3 := Address{Name: "Delhi"} 
    fmt.Println("Address3: ", a3) 
} 