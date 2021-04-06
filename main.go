package main 

import "fmt"

func ddate(a *int) {
fmt.Println(a)
*a = *a + 24
}
 
func main() {
var x int
var y *int
fmt.Println(x, &x, y, &y)
x = 10   //assingment
y = &x   //referencing
fmt.Println(x)//accessing
fmt.Println(y)
fmt.Println(*y)//dereferencing
ddate(&x)
fmt.Println(x)
}
