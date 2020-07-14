package main

import "fmt"

func main(){
	fmt.Println(commonSubstring("shishir", "madhura"))
}


func commonSubstring(source, target string) string{
	wordsDict := make(map[string]int)
	for _, value:= range source {
		wordsDict[string(value)] = 1
	}
	for _, value := range target{
		if _, ok := wordsDict[string(value)]; ok{
			return "Yes"
		}
	}
	return "No"
}