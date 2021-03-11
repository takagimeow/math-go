package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	list := []string{"1", "2", "3"}
	d := &DefaultUtils{list}
	got := d.Reverse()
	want := []string{"3", "2", "1"}

	for i, value := range got {
		if value != want[i] {
			t.Errorf("想定していない値が返されました")
		}
	}
}
