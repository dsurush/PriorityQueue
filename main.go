package main

import "fmt"

type list struct {
	firstElement *listNode
	lastElement  *listNode
	size  int
}

type listNode struct {
	prev     *listNode
	next     *listNode
	value    interface{}
	priority int
}

func (receiver *list) len() int {
	return receiver.size
}

func (receiver *list) first() interface{} {
	if receiver.firstElement == nil{
		return nil
	}
	return receiver.firstElement.value
}

func (receiver *list) last() interface{} {
	if receiver.lastElement == nil{
		return nil
	}
	return receiver.lastElement.value
}

func (receiver *list) equeue(value interface{}, priority int) {

	if receiver.len() == 0 {
		current := &listNode{
			prev:     nil,
			next:     nil,
			value:    value,
			priority: priority,
		}
		receiver.firstElement = current
		receiver.lastElement = current
		receiver.size++
		return
	}
	lastElement := receiver.lastElement
	current := &listNode{
		prev:     lastElement,
		next:     nil,
		value:    value,
		priority: priority,
	}
	lastElement.next = current
	receiver.lastElement = current
	receiver.size++
}

func (receiver *list) dequeue() interface{} {

	if receiver.len() == 0{
		return nil
	}

	if receiver.len() == 1{
		current := receiver.firstElement
		receiver.firstElement = nil
		receiver.lastElement = nil
		receiver.size--
		return current.value
	}

	receiver.size--
	current := receiver.firstElement
	receiver.firstElement = receiver.firstElement.next
	receiver.firstElement.prev = nil
	return current.value
}

func (receiver *list) PriorityQueue() {
	currentI := receiver.firstElement
	for ;currentI != nil; {
		currentJ := receiver.firstElement
		for ;currentJ.next != nil;{
			if currentJ.priority > currentJ.next.priority {
				fmt.Println(currentJ.priority, " > ", currentJ.next.priority)
				currentJ.priority, currentJ.next.priority = currentJ.next.priority, currentJ.priority
				currentJ.value, currentJ.next.value = currentJ.next.value, currentJ.value
			}
			currentJ = currentJ.next
		}
		currentI = currentI.next
	}
}
func main() {}