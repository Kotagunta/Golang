package main

import (
	"fmt"
)

type Node struct {
	Data any
	prev *Node
	next *Node
}

type LinkedList struct {
	Head *Node
	tail *Node
	len  int
}

func (a *LinkedList) insert(val any) {
	f := Node{Data: val}
	//f.prev=nil;
	//f.next=nil

	if a.len == 0 {
		a.Head = &f
		a.tail = &f
		a.len++

	} else {

		temp := a.tail
		temp.next = &f
		f.prev = temp
		a.tail = &f
		a.len++
	}

}

func (a *LinkedList) print() {
	temp := a.Head
	for i := 0; i < a.len; i++ {
		fmt.Println(temp.Data)
		temp = temp.next
	}
	fmt.Println()

}

func (a *LinkedList) printreverse() {
	temp := a.tail
	for i := 0; i < a.len; i++ {
		fmt.Println(temp.Data)
		temp = temp.prev
	}
	fmt.Println()
}

func (a *LinkedList) get(position any) *Node {
	c := a.Head
	for e := 0; e < a.len; e++ {
		if position == e {
			return c
		}
		c = c.next
	}
	return c
}

func (a *LinkedList) dekete(position any) {

}

//for i := 0; i<a.len; i++ {}

func main() {
	res := LinkedList{}
	res.insert(10)
	res.insert(20)
	res.insert(30)
	res.print()
	res.printreverse()
	f := res.get(0)
	fmt.Println(f.Data)
}
