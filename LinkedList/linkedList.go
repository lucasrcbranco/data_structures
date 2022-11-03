package LinkedList

import (
	"fmt"
	"github.com/lucasrcbranco/data_structures/Node"
)

type LinkedList struct {
	Head *Node.Node
	Tail *Node.Node
}

func (ll *LinkedList) Print() {
	pl := ll.Head
	for pl != nil {
		fmt.Println(pl.Value)
		pl = pl.Next
	}
}

func (ll *LinkedList) Insert(value int) {
	node := &Node.Node{Value: value}
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}

	ll.Tail.Next = node
	ll.Tail = node
}

func (ll *LinkedList) Remove(value int) {
	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
	}

	curr := ll.Head
	next := curr.Next
	for next != nil {
		if next.Value == value {
			curr.Next = next.Next
			break
		}
		curr = next
		next = curr.Next
	}
}
