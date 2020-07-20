package main

import (
	"fmt"
)

func main(){
	fmt.Println(findAnagrams("helo", "hi"))
	fmt.Println(findAnagrams("cde", "dcf"))
	fmt.Println(findAnagrams("cde", "abc"))
}

func findAnagrams(string1, string2 string) int{
		wordsDict := make(map[string]int)
		for i:= range string1 {
			wordsDict[string(string1[i])] += 1
		}
		for i:=range string2{
			wordsDict[string(string2[i])] += 1
		}
		counter := 0
		for _, value := range wordsDict{
			if (value % 2) != 0{
				counter ++
			}
		}
		return counter
}