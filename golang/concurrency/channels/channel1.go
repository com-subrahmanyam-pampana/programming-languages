package main

import "fmt"

func test1() {
	fmt.Println("<<<<<<<<<<<<<1st way of defining channel>>>>>>>>>>>..")
	//We use make() function to define a channel
	//Syntax var <your channle Name> chan <data type>
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T \n", a)
	}
}

func test2() {
	fmt.Println("<<<<<<<<<<<<<2nd way of defining channel>>>>>>>>>>>..")
	var a chan int = make(chan int)
	fmt.Println(a)
}

func test3() {
	fmt.Println("<<<<<<<<<<<<<Add data to channel>>>>>>>>>>>..")
	//Define the length of the channel as 5,or your own length
	var a chan int = make(chan int, 5)
	//Use to add data to channel
	fmt.Println(len(a))
	a <- 10
	a <- 12
	fmt.Println(len(a))
}

func test4() {
	fmt.Println("<<<<<<<<<<<<<Read data from channel>>>>>>>>>>>..")
	var a chan int = make(chan int, 5)
	a <- 10
	a <- 12
	d1 := <-a
	fmt.Println(d1)
	d2 := <-a
	fmt.Println(d2)
}
func main() {
	test1()
	test2()
	test3()
	test4()
}
