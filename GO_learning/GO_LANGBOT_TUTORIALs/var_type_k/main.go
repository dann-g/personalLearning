package main

import( 
	"fmt"
	"unsafe"
)
func main(){
	var age int //variable declaration
	fmt.Println("My age is", age)
	age = 29 //assignment
	fmt.Println("My age is", age)
	age = 54 //assignment
	fmt.Println("My new age is", age)
	
	//declaring a variable with an initialvalue
	var year int = 2022
	fmt.Println("The year is", year)
	//multiple variable declaration
	var width, height int = 100, 50
	fmt.Println("width is", width, "height is", height)
	
	//short hand declaration
	count:= 10
	fmt.Println("Count =", count)
	fmt.Println("\n")
	//bool type
	a:= true
	b:= false
	fmt.Println("a:", a,"b:", b)
	//int8,int16,int32,int64 blah
	y,x,c, d := 126, -32600, 2147483547, 2323232342334352
	fmt.Println("x is", x, "\ny is",y, "\nc is", c, "\nd is",d)
	fmt.Printf("y is %T and of size %d \n",y, unsafe.Sizeof(y))
	fmt.Printf("x is %T and of size %d \n",x, unsafe.Sizeof(x))
	fmt.Printf("c is %T and of size %d \n",c, unsafe.Sizeof(c))
	fmt.Printf("d is %T and of size %d \n",d, unsafe.Sizeof(d))
	//uint8, uint16,uint32, uint64
	var x0 uint8 = 255
	var x1 uint16 = 65535
	var x2 uint32 = 4294967295
	var x3 uint64 = 18446744073709551615
fmt.Printf(" %v x0 is %T and of size %d \n",x0,x0, unsafe.Sizeof(x0))
fmt.Printf("%v x1 is %T and of size %d \n",x1,x1, unsafe.Sizeof(x1))
fmt.Printf("%v x2 is %T and of size %d \n",x2,x2, unsafe.Sizeof(x2))
fmt.Printf("%v x3 is %T and of size %d \n",x3,x3, unsafe.Sizeof(x3))
//float32, float64 all are interpreed
 f1, f2 := 6.456456 , 23424.23423
 div := f2/f1
 fmt.Println("div of f1 and f2", div)
 //String type strings are a collection of bytes in Go
 first := "Daniel"
 last := "Gardner"
 name := first + " "+ last
 fmt.Println("My name is", name)
 
//constants and type
const trueConst= true
type myName string
const something myName = "Yeah"
fmt.Printf("\n%v something is %T and of size %d",something,something,unsafe.Sizeof(something))

}







