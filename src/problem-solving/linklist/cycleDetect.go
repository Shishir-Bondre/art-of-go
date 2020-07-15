package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func main(){
	head := &Node{}
	insertNewNodeAtBeginning(head, 1)
	insertNewNodeAtBeginning(head, 2)
	insertNewNodeAtBeginning(head, 3)
	insertNewNodeAtBeginning(head, 4)
	insertNewNodeAtBeginning(head, 5)
	traverseLinkedList(head)
	createCycleInLinkList(head, 3)
	detectCycle(head)
}

func insertNewNodeAtBeginning(head *Node, data int){
	node := Node{}
	node.data = data
	//fmt.Println(head.next, head.data)
	node.next = head.next
	head.next = &node
}

func traverseLinkedList(head *Node){
	node := head.next
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}


func createCycleInLinkList(head *Node, source int)*Node{
	node := head.next
	for node.data != source{
		node = node.next
	}
	sourceNode := node
	node = head.next // reset to start again to find target node now
	for node.next != nil{
		node = node.next
	}
	// make cycle
	node.next = sourceNode
	return head
}

func detectCycle(head *Node) bool{
	visited := make(map[int]int)
	node := head.next
	for node != nil{
		if value,ok:=visited[node.data];ok{
			fmt.Println("detected cycle at",value)
			return true
		}
		visited[node.data] = node.data
		node = node.next
	}
	fmt.Println("No cycle in this linked list")
	return false
}