package main

import "fmt"


type Node struct {
	Value int
	Next *Node
}


// 翻转单链表
func ReverseLinkList(link *Node) (*Node, error) {
	if link == nil {
		return nil, nil
	}
	
	var pre *Node
	current := link
	for (current != nil) {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}

	return pre, nil
}


func main() {
	linkList := &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, nil}}}}}}}
	cursor := linkList
	for cursor != nil {
		fmt.Printf("%d ", cursor.Value)
		cursor = cursor.Next
	}
	fmt.Printf("\n")

	reversedLink, _ := ReverseLinkList(linkList)
	cursor = reversedLink
	for cursor != nil {
		fmt.Printf("%d ", cursor.Value)
		cursor = cursor.Next
	}
	fmt.Printf("\n")
}
