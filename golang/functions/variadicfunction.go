package main

import "fmt"

/*
The function that is called with the varying number of arguments is
known as variadic function. Or in other words, a user is allowed to
 pass zero or more arguments in the variadic function
 function function_name(para1, para2...type)type {// code...}
*/

/*Veradic must be last paramters*/
func main() {
	//var sum int = addNNaturalNumbers(5, 1, 2, 3, 4, 5)
	//Or
	numbers := []int{1, 2, 3, 4, 5}
	var sum int = operationOnNNaturalNumbers("sum", numbers...)
	fmt.Println(sum)
}

func operationOnNNaturalNumbers(operationType string, x ...int) int {
	var sum int = 0

	for i := 0; i <= len(x); i++ {
		if operationType == "sum" {
			sum = sum + i
		}
		if operationType == "mul" {
			sum = sum * i
		}
	}
	return sum
}
