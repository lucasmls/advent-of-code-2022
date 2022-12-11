package main

import (
	"testing"
)

func Test_Marker(t *testing.T) {
	marker := NewMarker(4)

	marker.Add("a")
	marker.Add("b")
	marker.Add("c")
	marker.Add("d")
	marker.Add("e")
	marker.Add("f")

	got := marker.String()
	want := "cdef"

	if got != want {
		t.Errorf("Marker suite: = %v, want %v", got, want)
	}
}

func Test_MarkerValid(t *testing.T) {
	marker := NewMarker(4)

	marker.Add("a")
	marker.Add("b")
	marker.Add("c")
	marker.Add("d")
	marker.Add("e")
	marker.Add("f")

	got := marker.Valid()
	want := true
	if got != want {
		t.Errorf("Marker.Valid(): got = %v, want %v; data = %s", got, want, marker.String())
	}

	marker.Add("a")
	marker.Add("b")
	marker.Add("c")
	marker.Add("a")

	got = marker.Valid()
	want = false
	if got != want {
		t.Errorf("Marker.Valid(): got = %v, want %v; data = %s", got, want, marker.String())
	}

	m := NewMarker(4)
	str := "gzb"
	//str := "gzbzwzjwwr"
	for _, r := range str {
		char := string(r)
		m.Add(char)
	}

	got = m.Valid()
	want = false
	if got != want {
		t.Errorf("Marker.Valid(): got = %v, want %v; data = %s", got, want, marker.String())
	}
}
