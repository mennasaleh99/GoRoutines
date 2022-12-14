package main

import (
	"fmt"
	"time"
)

func normal() {
	fmt.Println("normal function")
}

func callnormalasgoroutines() {
	fmt.Println("callnormalasgoroutines")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, " : ", i)
		time.Sleep(time.Second)
	}
}

func main() {
	// sequential
	// fmt.Println("main () strat")
	// normal()
	// fmt.Println("main () end")
	// __________________________________________________________________________
	// concurrency
	// it will not print the body of method because
	// we need to apply delay to stop the main go routine which make our go routine stop working
	// fmt.Println("main () start 2")
	// go callnormalasgoroutines()
	// fmt.Println("main () end 2")
	// ____________________________________________________________________________________________
	// as time.sleep in function is one second and there is 3 iterations
	// we need to write 3 * or more , if we write 1* it will print only one iterations

	// fmt.Println("main () strat")
	// go f("go routines")
	// time.Sleep(3 * time.Second)
	// fmt.Println("main () end")

	// _______________________________________________________________________
	// anomynous function
	fmt.Println("main () strat")
	go func(msg string) {
		time.Sleep(2 * time.Second)
		fmt.Println(msg)
	}("Going")

	go func(msg string) {
		time.Sleep(time.Second)
		fmt.Println(msg)
	}("Going 1")

	go func(msg string) {

		fmt.Println(msg)
	}("Going 2")

	time.Sleep(3 * time.Second)
	fmt.Println("main () end")
}
