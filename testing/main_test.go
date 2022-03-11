package main

import "testing"

func TestSum(t *testing.T) {
	// total := Sum(5, 5)
	// if total != 10 {
	// 	t.Errorf("Sum was incorrect, got %d expect %d", total, 10)
	// }
	tables := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}
	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.c {
			t.Errorf("Sum was incorrect, got %d expect %d", total, item.c)
		}
	}
	res := GetMax(4, 6)
	if res != 6 {
		t.Errorf("GetMax was incorrect, got %d expect %d", res, 6)
	}
}
