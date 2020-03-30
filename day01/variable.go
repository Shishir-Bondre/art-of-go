package main

import "fmt"

// various variable declaration styles
func main() {
	// 1) standard syntax is var varName type; here no initialization
	var age int
	age = 10
	fmt.Println("The age is", age)
	// 2) if initialization is during declaration itself
	var age2 int = 20
	fmt.Println("This while declaration age: ", age2)
	// 3) Short hand declaration but the variable should be newly declared
	age3 := 30
	fmt.Println("This is the shorthand declaration of variables ", age3)
	// 4) No requirement of the type if initialization is there
	var age4 = 40
	fmt.Println("Type is not declared here age4: ", age4)
	// 5) Multiple declaration of variables
	var (
		age5 = 50
		name = "shishir"
		phone int
		)
	phone = 200
	fmt.Println("Multiple variables age : ", age5," name is: ", name, "phone is ", phone)

	// 6) For 5 one we can also have below syntax
	var age6, phone02 int
	var name02 string // Note here we can't declare multiple types in one line
	fmt.Println("Default values" ,age6, phone02, name02)
}