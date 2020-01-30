package main

import (
	"fmt"
	"testing"
)

func Test_list_equeue(t *testing.T) {
	q := list{}

	if q.len() != 0{
		want := 0
		got := q.len()
		t.Error("method len() in M(equeue) is not correct for empty queue got", got, "want", want)
	}

	if got := q.first(); got != nil{
		t.Error("method first() in M(equeue) is not correct for empty queue got", got, "want", nil)
	}

	if got := q.last(); got != nil{
		t.Error("method last() in M(equeue) is not correct for empty queue got", got, "want", nil)
	}

	q.equeue("Deda", 54)

	if q.len() != 1{
		want := 1
		got := q.len()
		t.Error("method len() in M(equeue) is not correct for 1 element in queue got", got, "want", want)
	}

	if got := q.first(); got != "Deda"{
		want := "Deda"
		t.Error("method first() in M(equeue) is not correct for 1 element in queue got", got, "want", want)
	}

	if got := q.last(); got != "Deda"{
		want := "Deda"
		t.Error("method last() in M(equeue) is not correct for 1 element in queue got", got, "want", want)
	}

	q.equeue("Babushka", 22)
	q.equeue("Vnuchka", 16)

	if q.len() != 3{
		want := 3
		got := q.len()
		t.Error("method len() in M(equeue) is not correct for queue which have many element: got", got, "want", want)
	}

	if got := q.first(); got != "Deda"{
		want := "Deda"
		t.Error("method first() in M(equeue) is not correct for queue which have many element: got", got, "want", want)
	}

	if got := q.last(); got != "Vnuchka"{
		want := "Vnuchka"
		t.Error("method last() in M(equeue) is not correct for queue which have many element: got", got, "want", want)
	}
}

func Test_list_dequeue(t *testing.T) {
	q := list{}
	q.equeue("Babushka", 22)
	q.equeue("Vnuchka", 16)

	deleteElement := q.dequeue()

	if q.len() != 1{
		want := 1
		got := q.len()
		t.Error("method len() in M(dequeue) is not correct for queue which have many element: got", got, "want", want)
	}

	if got := q.first(); got != "Vnuchka"{
		want := "Vnuchka"
		t.Error("method first() in M(dequeue) is not correct for queue which have many element: got", got, "want", want)
	}

	if got := q.last(); got != "Vnuchka"{
		want := "Vnuchka"
		t.Error("method last() in M(dequeue) is not correct for queue which have many element: got", got, "want", want)
	}

	if deleteElement != "Babushka"{
		want := "Babushka"
		t.Error("method (dequeue) is not correct for queue which have many element: DeleteElement", deleteElement, "want", want)
	}

	deleteElement = q.dequeue()

	if q.len() != 0{
		want := 0
		got := q.len()
		t.Error("method len() in M(dequeue) is not correct for 1 element in queue got", got, "want", want)
	}

	if got := q.first(); got != nil{
		t.Error("method first() in M(dequeue) is not correct for 1 element in queue got", got, "want", nil)
	}

	if got := q.last(); got != nil{
		t.Error("method last() in M(dequeue) is not correct for 1 element in queue got", got, "want", nil)
	}

	if deleteElement != "Vnuchka"{
		want := "Vnuchka"
		t.Error("method (dequeue) is not correct for 1 element in queue element: DeleteElement", deleteElement, "want", want)
	}

	deleteElement = q.dequeue()

	if q.len() != 0{
		want := 0
		got := q.len()
		t.Error("method len() in M(dequeue) is not correct for empty queue got", got, "want", want)
	}

	if got := q.first(); got != nil{
		t.Error("method first() in M(dequeue) is not correct for empty queue got", got, "want", nil)
	}

	if got := q.last(); got != nil{
		t.Error("method last() in M(dequeue) is not correct for 1 empty queue got", got, "want", nil)
	}

	if deleteElement != nil{
		t.Error("method (dequeue) is not correct for empty queue element: DeleteElement", deleteElement, "want", nil)
	}
}

var q = list{}

func ExamplePrioritizeQueueForEmptyQueue() {
	q.PrioritizeQueue()
	fmt.Println(q)
	// Output: {<nil> <nil> 0}
}

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