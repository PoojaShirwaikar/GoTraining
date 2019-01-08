package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(5, 5)
	if total != 10 {
		t.Errorf("sum was correct, got %d want %d", total, 10)
	}

}
