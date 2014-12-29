package containers

import (
	"testing"
)

func TestAppend(t *testing.T) {
	arr := NewArrayList(10)
	expected := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr.Append(i)
		expected[i] = i
	}
	for i:=0; i < 10; i++ {
		if(arr.Get(i).(int) != expected[i]) {
			t.Errorf("Error at Append %d\n", i)
		}
	}
}

func TestPrepend(t *testing.T) {
	arr := NewArrayList(10)
	expected := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr.Prepend(i)
		expected[9-i] = i
	}
	for i:=0; i < 10; i++ {
		t.Log(arr.Get(i).(int))
		if(arr.Get(i).(int) != expected[i]) {
			t.Errorf("Error at Prepend %d\n", i)
		}
	}
}

func TestInsert(t *testing.T) {
	arr := NewArrayList(10)
	expected := make([]int, 10)
	for i := 0; i < 10; i++ {
		if i != 0 && i != 2 && i !=3 && i != 9 {
			arr.Append(i)
		}
		expected[i] = i
	}
	arr.Insert(0, 0)
	arr.Insert(2, 2)
	arr.Insert(3, 3)
	arr.Insert(9, 9)
	for i:=0; i < 10; i++ {
		t.Log(arr.Get(i).(int))
		if(arr.Get(i).(int) != expected[i]) {
			t.Errorf("Error at Insert %d\n", i)
		}
	}
}
