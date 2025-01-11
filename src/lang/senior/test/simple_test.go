package test

import "testing"

func TestAdd(t *testing.T) {
	a := 1
	b := 2
	var p = 4

	r := Add(a, b)
	if r != p {
		t.Error("和预期结果不一致")
	}
}
