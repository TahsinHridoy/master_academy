
package main

import "fmt"
import "reflect"

func main() {
//primitive datatype

//rune byte = alias
// string, int, float32, bool
//composite datatype


//slice

var students [3]string 
/*
students[0]="asgor"
students[1]="nasarul"
students[2]="jahid"
z := students[0:2]
*/
//x :=make([]string, 6)
var fruits []int
fruits = append(fruits, 34, 65, 75)
//map
//struct

fmt.Println(fruits)
fmt.Printf("%T\n", fruits)
fmt.Printf("%T\n", students)
b := reflect.TypeOf(fruits).Kind().String()
fmt.Println(b)

}
