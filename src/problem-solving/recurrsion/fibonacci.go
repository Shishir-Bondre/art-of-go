package main

import "fmt"


func main(){
	cache := make(map[int]int)
	cache[0] = 0
	cache[1] = 1
	fmt.Println(fibonacciNos(6, &cache))
}

func fibonacciNos(no int, cache *map[int]int)int{
	if value, ok := (*cache)[no]; ok{
		return value
	}
	(*cache)[no] = fibonacciNos(no-1, cache) + fibonacciNos(no-2, cache)
	return (*cache)[no]
}