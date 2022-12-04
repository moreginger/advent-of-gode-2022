package main

import (
	"testing"
)

func Test(t *testing.T) {
	l := Range{
		min: 2,
		max: 62,
	}
	r := Range{
		min: 62,
		max: 98,
	}
	if !l.Overlaps(r) {
		t.Errorf("Failed")
	}
	if !r.Overlaps(l) {
		t.Errorf("Failed")
	}
}
