package main

import (
	"fmt"
)

func floatToIntConversion() {
	/*
		int + float64 not allowed.
		So convery var "j" to int
	*/
	i := 55
	j := 67.5
	//sum := i + j //not allowed
	sum := i + int(j)
	fmt.Println(sum)
}

func intToFloatConversion() {
	fmt.Println("<<<<<<<Method 2 >>>>>>>>>")
	i := 55
	j := 67.5
	var k float64 = float64(i)
	fmt.Println(i)
	sum2 := j + k
	fmt.Println(sum2)
}

func method3() {
	// character to ASCII
	char := 'a' // rune, not string
	ascii := int(char)
	fmt.Println(string(char), " : ", ascii)

	// ASCII to character

	asciiNum := 65 // Uppercase A
	character := string(rune(asciiNum))
	fmt.Println(asciiNum, " : ", character)

}
func typeConversion() {
	floatToIntConversion()
	intToFloatConversion()
}
