package main

import (
	"reflect"
	"testing"
)

func Test_calcLetterPriority(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Lowercase
		{
			name: "a",
			args: args{
				item: "a",
			},
			want: 1,
		},
		{
			name: "b",
			args: args{
				item: "b",
			},
			want: 2,
		},
		{
			name: "c",
			args: args{
				item: "c",
			},
			want: 3,
		},
		{
			name: "x",
			args: args{
				item: "x",
			},
			want: 24,
		},
		{
			name: "y",
			args: args{
				item: "y",
			},
			want: 25,
		},
		{
			name: "z",
			args: args{
				item: "z",
			},
			want: 26,
		},

		// Uppercase
		{
			name: "A",
			args: args{
				item: "A",
			},
			want: 27,
		},
		{
			name: "B",
			args: args{
				item: "B",
			},
			want: 28,
		},
		{
			name: "C",
			args: args{
				item: "C",
			},
			want: 29,
		},
		{
			name: "X",
			args: args{
				item: "X",
			},
			want: 50,
		},
		{
			name: "Y",
			args: args{
				item: "Y",
			},
			want: 51,
		},
		{
			name: "Z",
			args: args{
				item: "Z",
			},
			want: 52,
		},

		// Examples
		{
			name: "p",
			args: args{
				item: "p",
			},
			want: 16,
		},
		{
			name: "L",
			args: args{
				item: "L",
			},
			want: 38,
		},
		{
			name: "P",
			args: args{
				item: "P",
			},
			want: 42,
		},
		{
			name: "v",
			args: args{
				item: "v",
			},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcLetterPriority(tt.args.item); got != tt.want {
				t.Errorf("itemToPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chunk(t *testing.T) {
	type args struct {
		slice []string
		size  int
	}

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			args: args{
				slice: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
				size:  3,
			},
			want: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
		},
		{
			args: args{
				slice: []string{"1", "2", "3", "4", "5", "6", "7", "8"},
				size:  2,
			},
			want: [][]string{
				{"1", "2"},
				{"3", "4"},
				{"5", "6"},
				{"7", "8"},
			},
		},
		{
			args: args{
				slice: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
				size:  2,
			},
			want: [][]string{
				{"1", "2"},
				{"3", "4"},
				{"5", "6"},
				{"7", "8"},
				{"9"},
			},
		},
		{
			args: args{
				slice: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
				size:  4,
			},
			want: [][]string{
				{"1", "2", "3", "4"},
				{"5", "6", "7", "8"},
				{"9"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chunk(tt.args.slice, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
