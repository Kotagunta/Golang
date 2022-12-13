package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	Head *Node
	len  int
}

func (a *LinkedList) insert(e int) {
	g := Node{value: e}

	g.next = nil
	if a.len == 0 {
		a.Head = &g
		a.len += 1
		fmt.Println(a.Head.value)

	} else {
		for w := 0; w < a.len; w++ {
			if a.Head.next == nil {
				a.Head.next = &g
				a.len++
			}
			if w == a.len-1 {
				fmt.Println(a.Head.value)
			}

			a.Head = a.Head.next

		}

	}

}

func main() {
	var siva = LinkedList{}
	siva.insert(10)
	siva.insert(1000)

}
