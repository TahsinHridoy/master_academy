package main 

import "fmt"
/*
//example_1
func add(x int, y int)int {
r := x + y
return r
}

//example_2
func add(x, y int)int {
r := x + y
return r
}

//example_3
func add(x, y int) (r int) {
r = x + y
return r
}


//example_4
func add(x, y int)(r int) {
r = x + y
return
}
 
func rec(k, b int) (c, a int){
c = 2 * (k+b)
a = k*b
return
}

func update(a *int, b *string) {
*a = *a + 5      //derefarencing
*b = *b +"doe"
return
}
*/
func main() {
/*
//e, f := rec(12, 28)
x := 9
y := "herry "
update(&x, &y)  //refarencing
*/
a := func(x, y int)(r int){
r = x*y
return
}(19, 10)
fmt.Println(a)
}