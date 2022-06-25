package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		target coord
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,1]",
			args: args{
				target: coord{row: 1, col: 1},
			},
			want: 20151125,
		},

		{
			name: "[1,2]",
			args: args{
				target: coord{row: 1, col: 2},
			},
			want: 18749137,
		},

		{
			name: "[2,1]",
			args: args{
				target: coord{row: 2, col: 1},
			},
			want: 31916031,
		},
		{
			name: "[4,3]",
			args: args{
				target: coord{row: 4, col: 3},
			},
			want: 21345942,
		},
		{
			name: "[6,6]",
			args: args{
				target: coord{row: 6, col: 6},
			},
			want: 27995004,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.target); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextCoord(t *testing.T) {
	tests := []struct {
		name    string
		current coord
		want    coord
	}{
		{
			name:    "[1,2]",
			current: coord{row: 1, col: 2},
			want:    coord{row: 3, col: 1},
		},
		{
			name:    "[4,3]",
			current: coord{row: 4, col: 3},
			want:    coord{row: 3, col: 4},
		},
		{
			name:    "[6,6]",
			current: coord{row: 6, col: 6},
			want:    coord{row: 5, col: 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextCoord(tt.current); got.col != tt.want.col && got.row != tt.want.row {
				t.Errorf("nextCoord() = %v, want %v", got, tt.want)
			}
		})
	}
}
