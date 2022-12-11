package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

type Marker struct {
	cap  int
	data []string
}

func NewMarker(capacity int) *Marker {
	return &Marker{
		cap:  capacity,
		data: []string{},
	}
}

func (s *Marker) Add(value string) {
	if s.Size() < s.cap {
		s.data = append(s.data, value)
		return
	}

	s.rotate()
	s.data[s.cap-1] = value
}

func (s *Marker) Size() int {
	return len(s.data)
}

func (s *Marker) String() string {
	return strings.Join(s.data, "")
}

func (s *Marker) Valid() bool {
	if len(s.data) != s.cap {
		return false
	}

	set := map[string]bool{}
	for _, v := range s.data {
		set[v] = true
	}

	if len(set) == s.cap {
		return true
	}

	return false
}

func (s *Marker) rotate() {
	for i := 1; i < s.cap; i++ {
		s.data[i-1] = s.data[i]
	}
}

func main() {
	marker := NewMarker(14)

	for i, r := range input {
		char := string(r)
		marker.Add(char)

		if marker.Valid() {
			fmt.Println("Result", i+1)
			break
		}
	}
}
