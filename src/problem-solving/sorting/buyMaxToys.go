package main

import (
	"fmt"
	"sort"
)

func main(){
	maximumToys([]int{1, 2, 3, 4}, 6)
}

func maximumToys(prices []int, k int){
	sort.Slice(prices, func(i, j int)bool{ return prices[i] < prices[j]})
	noToysCanBePurchase := 0
	amountLeft := k
	for _,value:= range prices{
		if amountLeft > 0 {
			amountLeft -= value
			noToysCanBePurchase ++
		}
	}
	fmt.Println(noToysCanBePurchase)
}
