package set

import (
	"testing"
)

type myStruct struct {
	val1 int
	val2 bool
}

func TestSet_WithStruct(t *testing.T) {
	set := NewHashSet[myStruct]()

	s1 := myStruct{
		val1: 1,
		val2: false,
	}

	s2 := myStruct{
		val1: 2,
		val2: true,
	}

	set.Add(s1)
	set.Add(s2)

	size := set.Size()
	if size != 2 {
		t.Errorf("expected size to be %d, but is %d", 2, size)
	}

	contains := set.Contains(s1)
	if contains == false {
		t.Error("expected contains to be true, but was false")
	}

	contains = set.Contains(s2)
	if contains == false {
		t.Error("expected contains to be true, but was false")
	}

	set.Remove(s1)
	size = set.Size()
	if size != 1 {
		t.Errorf("expected size to be %d, but is %d", 1, size)
	}

	contains = set.Contains(s1)
	if contains == true {
		t.Error("expected contains to be false, but was true")
	}

	contains = set.Contains(s2)
	if contains == false {
		t.Error("expected contains to be true, but was false")
	}

	set.Remove(s2)
	size = set.Size()
	if size != 0 {
		t.Errorf("expected size to be %d, but is %d", 0, size)
	}

	contains = set.Contains(s1)
	if contains == true {
		t.Error("expected contains to be false, but was true")
	}

	contains = set.Contains(s2)
	if contains == true {
		t.Error("expected contains to be false, but was true")
	}
}

func TestSet(t *testing.T) {
	set := NewHashSet[int]()

	var num1 int = 1
	var num2 int = 2
	var num3 int = 3

	set.Add(1)
	set.Add(2)
	set.Add(3)

	size := set.Size()
	if size != 3 {
		t.Errorf("expected size to be %d, but is %d", 3, size)
	}

	contains := set.Contains(num1)
	if contains == false {
		t.Error("expected contains to be true, but was false")
	}

	contains = set.Contains(num2)
	if contains == false {
		t.Error("expected contains to be true, but was false")
	}

	contains = set.Contains(num3)
	if contains == false {
		t.Error("expected contains to be true, but was false")
	}

	set.Remove(num1)
	size = set.Size()
	if size != 2 {
		t.Errorf("expected size to be %d, but is %d", 2, size)
	}

	contains = set.Contains(num1)
	if contains == true {
		t.Error("expected contains to be false, but was true")
	}

	contains = set.Contains(num2)
	if contains == false {
		t.Error("expected contains to be true, but was false")
	}

	contains = set.Contains(num3)
	if contains == false {
		t.Error("expected contains to be true, but was false")
	}

	set.Remove(num2)
	size = set.Size()
	if size != 1 {
		t.Errorf("expected size to be %d, but is %d", 1, size)
	}

	contains = set.Contains(num2)
	if contains == true {
		t.Error("expected contains to be false, but was true")
	}

	contains = set.Contains(num3)
	if contains == false {
		t.Error("expected contains to be true, but was false")
	}

	set.Remove(num3)
	size = set.Size()
	if size != 0 {
		t.Errorf("expected size to be %d, but is %d", 0, size)
	}

	contains = set.Contains(num3)
	if contains == true {
		t.Error("expected contains to be false, but was true")
	}
}

func TestRemove_WithNoElements(t *testing.T) {
	set := NewHashSet[int]()

	set.Remove(1)

	// Assert no error
}
