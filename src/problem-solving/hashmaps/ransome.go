package main

import "fmt"

func main(){
	magazine := []string {"two", "times", "is", "not", "four"}
	note := []string {"two", "times", "is", "two", "four"}
	fmt.Println(checkMagazine(magazine, note))
}

func checkMagazine(magazine []string, note []string) string {

	magazineDict := make(map[string]int)
	for _, value:=range magazine {
		magazineDict[value] = 1
	}
	present := true
	for _,value := range note{
		if _, ok := magazineDict[value]; ok==false{
			present = false
			break
		}else{
			delete(magazineDict, value)
		}
	}
	if present {
		return "Yes"
	}
	return "No"
}
