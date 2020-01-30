package main

import "fmt"

var q = list{}

func ExamplePrioritizeQueue() {
	q.equeue("Babushka", 54)
	q.equeue("Dedushka", 34)
	q.equeue("Vnuchka", 12)
	q.PrioritizeQueue()
	current := q.firstElement
	for ; current != nil ; {
		fmt.Print(current.value, " ", current.priority, " ")
		current = current.next
	}
	// Output: Vnuchka 12 Dedushka 34 Babushka 54
}
