package main 

import( "fmt"
  "learnpackage/simpleinterest"
  "log"
 )
 var principal = -5000.0
 var rate = 10.0
 var time = 1.0
func init() {
fmt.Println("Main package initialized")
if principal < 0{
	log.Fatal("Principal is less than zero")
}
if rate < 0{
	log.Fatal("Rate of interest is less than zero")
}
if time < 0{
	log.Fatal("Duration is less than zero")
 }
}
func main(){
 fmt.Println("Simple Interest Calculation")
 interest_est:= simpleinterest.Calculate(principal,rate,time)
 fmt.Println("Simple interest is",interest_est)
}
