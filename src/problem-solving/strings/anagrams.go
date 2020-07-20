package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	fmt.Println(findAnagrams("helo", "hi"))
	fmt.Println(findAnagrams("cde", "dcf"))
	fmt.Println(findAnagrams("cde", "abc"))
}

func findAnagrams(input1, input2 string) int{
		string1 := strings.Split(input1, "")
		string2 := strings.Split(input2, "")
		sort.Strings(string1)
		sort.Strings(string2)
		fmt.Println(string1, string2)
		smallestLength := len(string2)
		largestLength := len(string1)
		if largestLength < smallestLength {
			smallestLength, largestLength = largestLength, smallestLength
		}
		counter := 0
		for i:=0 ; i < smallestLength; i++{
			if string1[i] != string2[i]{
				counter ++
			}
		}
		counter += largestLength - smallestLength
		return counter
}