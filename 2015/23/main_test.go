package main

import (
	"testing"
)

func Test_performOperations(t *testing.T) {
	type args struct {
		instructions []instruction
	}
	tests := []struct {
		name  string
		args  args
		want  uint64
		want1 uint64
	}{
		{
			name: "Simple case",
			args: args{
				instructions: []instruction{
					{name: "inc", value: "a"},
					{name: "jio", value: "a", third: +2},
					{name: "tpl", value: "a"},
					{name: "inc", value: "a"},
				},
			},
			want:  2,
			want1: 0,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := performOperations(0, 0, tt.args.instructions)
			if got != tt.want {
				t.Errorf("performOperations() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("performOperations() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Help"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			part1()
		})
	}
}
