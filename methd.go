package main 

import "fmt"

type Email interface{
write(string, string, string)
send()
read()


}

type test struct {
x string
y string
z int
}
func (r test) write(to string, from string, sub int){
fmt.Println(r.x, to, from, r.y)
}

func (r test) send(){
fmt.Println(r.y)
}
func (r test) read(){
fmt.Println(r.z)
}
type doctor struct {
Name string 
Age int 
Salary float32
}
//method
func (d doctor)speak() {
fmt.Println(d.Name,"can speak")
}

func (d doctor)getName()string {
return d.Name
}

func (d doctor)getSelaryInfo()float32{
return d.Salary
}
func main() {
//var d=doctor{ "tareq", 32, 5000 }
//var d doctor{ Name:"tareq", Age:32, Salary:5000, }
/*
var d doctor 
d.Name= "tareq"
d.Age= 34
d.Salary= 5000

fmt.Println(d.getName())
fmt.Println(d.getSelaryInfo())
*/
//var e Email
var d test
d.x = "te@gmail"
d.y = "hi@gmail"
d.z = 67
d.write(d.x, d.y, d.z)
d.send()
d.read()
}





