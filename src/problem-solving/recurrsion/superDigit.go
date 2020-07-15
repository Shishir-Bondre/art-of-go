package main

import "fmt"

func main(){
	fmt.Println(superDigit(9875))
	//fmt.Println(9/10 )
}

func superDigit(no int)int{
	noDigits := 0
	sum := 0
	for no > 0{
		sum += no % 10
		no = no / 10
		//fmt.Println(no, sum, noDigits)
		noDigits ++
	}
	if noDigits > 1{
		return superDigit(sum)
	}
	return sum
}