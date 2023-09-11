package main

import "fmt"

func printType(data interface{}) {
	fmt.Printf("%v is %T\n", data, data)
}

func main() {
	printType("str")
	printType(228)
	printType(13.37)
	printType(make(chan int))
}
