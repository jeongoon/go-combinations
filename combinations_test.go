package  combinations_test

import "testing"
import "reflect"

import combo "github.com/jeongoon/go-combinations"

func TestSomeIndexes_NumberOfCases(t *testing.T) {
	res1 := combo.SomeIndexes(7, 5)
	res1_len := len(res1)

	if res1_len != 21 {
		t.Fatalf("expected %d result but got %d", 21, res1_len)
	}
}

func TestAllIndexes_StrictOrder(t *testing.T) {
	res1 := combo.AllIndexes(5)
	res1_expected := [][]int{ {0},{1},{2},{3},{4},
		{0,1},{0,2},{0,3},{0,4},
		{1,2},{1,3},{1,4},{2,3},{2,4},{3,4},
		{0,1,2},{0,1,3},{0,1,4},{0,2,3},{0,2,4},{0,3,4},{1,2,3},{1,2,4},{1,3,4},{2,3,4},
		{0,1,2,3},{0,1,2,4},{0,1,3,4},{0,2,3,4},{1,2,3,4},{0,1,2,3,4} }

	if ! reflect.DeepEqual(res1, res1_expected) {
		t.Fatalf("expcedd order is %v but got %v", res1_expected, res1)
	}
}
