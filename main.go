package main

import "github.com/lucasrcbranco/data_structures/LinkedList"

func main() {
	ll := LinkedList.LinkedList{}
	ll.Insert(80)
	ll.Insert(10)
	ll.Insert(50)
	ll.Insert(20)
	ll.Insert(60)
	ll.Insert(30)
	ll.Insert(40)
	ll.Insert(70)

	ll.Remove(70)
	ll.Remove(20)
	ll.Print()
}
