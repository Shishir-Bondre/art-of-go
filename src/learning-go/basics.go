package main

import (
	"fmt"
	"runtime"
)

func main(){
	// call the square root function
	fmt.Println(squareRoot(9))
	switchCase(2)
	printOS()
}

func switchCase(exp int) {
	switch exp {
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
	default:
		fmt.Println("Default case")
	}
}

func printOS() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}
}

func squareRoot(x int) int {
	var z int
	for i:=1; i*i <= x; i++{
		z = i
	}
	return z
}