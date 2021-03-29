package main 
import "fmt"


func main() {
fmt.Println("input  your name and age:")
var name string
var age int

fmt.Scanf("%s %d", &name, &age)
fmt.Printf("your name %s & your age is %d", name, age)

}
