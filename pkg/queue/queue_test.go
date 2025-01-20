package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	var queue = New[*int]()
	var num1 int = 1
	var num2 int = 2
	var num3 int = 3

	queue.Enqueue(&num1)
	queue.Enqueue(&num2)
	queue.Enqueue(&num3)

	actualNum1, removed := queue.Dequeue()
	if removed == false {
		t.Error("expected removed to be true, but was false")
	}

	if actualNum1 != &num1 {
		t.Errorf("expected %d, but got %d", num1, actualNum1)
	}

	actualNum2, removed := queue.Dequeue()
	if removed == false {
		t.Error("expected removed to be true, but was false")
	}

	if actualNum2 != &num2 {
		t.Errorf("expected %d, but got %d", num2, actualNum2)
	}

	actualNum3, removed := queue.Dequeue()
	if removed == false {
		t.Error("expected removed to be true, but was false")
	}

	if actualNum3 != &num3 {
		t.Errorf("expected %d, but got %d", num3, actualNum3)
	}
}

func TestQueue_Peek(t *testing.T) {
	var queue = New[*int]()
	var num1 int = 1
	var num2 int = 2
	var num3 int = 3

	queue.Enqueue(&num1)
	queue.Enqueue(&num2)
	queue.Enqueue(&num3)

	actualNum1, exists := queue.Peek()
	if exists == false {
		t.Error("expected exists to be true, but was false")
	}

	if num1 != *actualNum1 {
		t.Errorf("expected %d, but got %d", num1, actualNum1)
	}

	queue.Dequeue()

	actualNum2, exists := queue.Peek()
	if exists == false {
		t.Error("expected exists to be true, but was false")
	}

	if num2 != *actualNum2 {
		t.Errorf("expected %d, but got %d", num2, actualNum2)
	}

	queue.Dequeue()

	actualNum3, exists := queue.Peek()
	if exists == false {
		t.Error("expected exists to be true, but was false")
	}

	if num3 != *actualNum3 {
		t.Errorf("expected %d, but got %d", num3, actualNum3)
	}

	queue.Dequeue()

	// After the last element has been removed
	actualNum4, exists := queue.Peek()
	if exists == true {
		t.Error("expected exists to be false, but was true")

	}

	if actualNum4 != nil {
		t.Error("expected actualNum4 to be nil")
	}
}
func TestQueue_Size(t *testing.T) {
	var queue = New[*int]()
	var num1 int = 1
	var num2 int = 2
	var num3 int = 3

	queue.Enqueue(&num1)
	queue.Enqueue(&num2)
	queue.Enqueue(&num3)

	if queue.Size() != 3 {
		t.Errorf("expected size %d, but got %d", 3, queue.Size())
	}

	queue.Dequeue()
	if queue.Size() != 2 {
		t.Errorf("expected size %d, but got %d", 2, queue.Size())
	}

	queue.Dequeue()
	if queue.Size() != 1 {
		t.Errorf("expected size %d, but got %d", 1, queue.Size())
	}

	queue.Dequeue()
	if queue.Size() != 0 {
		t.Errorf("expected size %d, but got %d", 0, queue.Size())
	}
}

func TestQueue_RemoveWithNoElements(t *testing.T) {
	var queue = New[*int]()

	actual, removed := queue.Dequeue()
	if removed == true {
		t.Error("expected removed to be true, but was false")
	}

	if actual != nil {
		t.Errorf("expected %d, to be nil", actual)
	}
}
