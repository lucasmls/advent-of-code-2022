package main

import "testing"

func TestAssignment_FullyOverlap(t *testing.T) {
	type fields struct {
		Raw string
	}
	type args struct {
		counterpartyAssignment Assignment
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			fields: fields{
				Raw: "38-41",
			},
			args: args{
				counterpartyAssignment: NewAssignment("38-38"),
			},
			want: true,
		},
		{
			fields: fields{
				Raw: "10-20",
			},
			args: args{
				counterpartyAssignment: NewAssignment("12-18"),
			},
			want: true,
		},
		{
			fields: fields{
				Raw: "10-20",
			},
			args: args{
				counterpartyAssignment: NewAssignment("20-25"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAssignment(tt.fields.Raw)

			if got := a.FullyOverlap(tt.args.counterpartyAssignment); got != tt.want {
				t.Errorf("FullyOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssignment_PartiallyOverlap(t *testing.T) {
	type fields struct {
		Raw string
	}
	type args struct {
		counterpartyAssignment Assignment
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			fields: fields{
				"5-7",
			},
			args: args{
				counterpartyAssignment: NewAssignment("7-9"),
			},
			want: true,
		},
		{
			fields: fields{
				"2-8",
			},
			args: args{
				counterpartyAssignment: NewAssignment("3-7"),
			},
			want: true,
		},
		{
			fields: fields{
				"6-6",
			},
			args: args{
				counterpartyAssignment: NewAssignment("4-6"),
			},
			want: true,
		},
		{
			fields: fields{
				"2-6",
			},
			args: args{
				counterpartyAssignment: NewAssignment("4-8"),
			},
			want: true,
		},
		{
			fields: fields{
				"2-6",
			},
			args: args{
				counterpartyAssignment: NewAssignment("7-10"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAssignment(tt.fields.Raw)

			if got := a.PartiallyOverlap(tt.args.counterpartyAssignment); got != tt.want {
				t.Errorf("PartiallyOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
