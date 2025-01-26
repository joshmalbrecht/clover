package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	var stack = New[int]()

	var num1 int = 1
	var num2 int = 2
	var num3 int = 3

	stack.Push(num1)
	stack.Push(num2)
	stack.Push(num3)

	size := stack.Size()
	if size != 3 {
		t.Errorf("expected size to be %d, but is %d", 3, size)
	}

	actual, removed := stack.Pop()
	if removed == false {
		t.Error("expected removed to be true, but was false")
	}

	if actual != num3 {
		t.Errorf("expected %d, but was %d", num3, actual)
	}

	size = stack.Size()
	if size != 2 {
		t.Errorf("expected size to be %d, but is %d", 2, size)
	}

	actual, removed = stack.Pop()
	if removed == false {
		t.Error("expected removed to be true, but was false")
	}

	if actual != num2 {
		t.Errorf("expected %d, but was %d", num2, actual)
	}

	size = stack.Size()
	if size != 1 {
		t.Errorf("expected size to be %d, but is %d", 1, size)
	}

	actual, removed = stack.Pop()
	if removed == false {
		t.Error("expected removed to be true, but was false")
	}

	if actual != num1 {
		t.Errorf("expected %d, but was %d", num1, actual)
	}

	size = stack.Size()
	if size != 0 {
		t.Errorf("expected size to be %d, but is %d", 0, size)
	}
}

func TestPeek(t *testing.T) {
	var stack = New[int]()

	var num1 int = 1
	var num2 int = 2

	stack.Push(num1)
	stack.Push(num2)

	actual, existing := stack.Peek()
	if !existing {
		t.Error("expected existing to be true, but was false")
	}

	if actual != num2 {
		t.Errorf("expected %d, but was %d", num2, actual)
	}

	stack.Pop()

	actual, existing = stack.Peek()
	if !existing {
		t.Error("expected existing to be true, but was false")
	}

	if actual != num1 {
		t.Errorf("expected %d, but was %d", num1, actual)
	}
}

func TestPop_WithNoElements(t *testing.T) {
	var stack = New[int]()

	_, removed := stack.Pop()
	if removed {
		t.Error("expected removed to be false, but was true")
	}
}

func TestPeek_WithNoElements(t *testing.T) {
	var stack = New[int]()

	_, existing := stack.Peek()
	if existing {
		t.Error("expected existing to be false, but was true")
	}
}
