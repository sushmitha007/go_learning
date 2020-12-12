package main

import "fmt"

type square struct{
	sideLength float64
}
type triangle struct{
	height float64
	base float64
}
type shape interface{
	getArea() float64
}
 func main(){
	sq:=square{6.7}
	tr:=triangle{
		5,
		3,
	}
printArea(sq)
printArea(tr)
 }
 func(t triangle) getArea() float64{
	result:= 0.5*t.base* t.height
     os.Exit(1)
	return result
 }

 func(s square) getArea() float64{
	 return s.sideLength*s.sideLength
 }
 func printArea(s shape){
  fmt.Println(s.getArea())
}