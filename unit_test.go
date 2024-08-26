package main

import "testing"

func TestName(t *testing.T) {
	r := get_randam(100)
	if r < 0 || r > 100 {
		t.Errorf("random not in 0~100")
	}
}
