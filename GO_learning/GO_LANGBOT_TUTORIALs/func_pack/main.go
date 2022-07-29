package main

import "fmt"
//if of same type you don't need to specifiy type
func calculateBill(price , no int) int {
	var totalPrice = price * no
	return totalPrice
}
//named return values don't need to be called with return keyword
func rectProps(length, width float64)(area,perimeter float64){
	var area = length *width
	var perimeter = (length + width) *2
	return 
}
func main(){
	price, no := 90, 6
	totalPrice := calculateBill(price,no)
	fmt.Println("Total price is", totalPrice)
	area, perimeter := rectProps(10.6, 2323.54563)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)
	//running go install with func args
}


