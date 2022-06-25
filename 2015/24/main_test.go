package main

import (
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{

			name: "Example",
			args: args{
				[]int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
			},
			want: 99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{

			name: "Example",
			args: args{
				[]int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
			},
			want: 44,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_sumToK(t *testing.T) {
	type args struct {
		numbers []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1-5 sum to 4",
			args: args{
				numbers: []int{1, 2, 3, 4, 5},
				k:       4,
			},
			want: [][]int{{1, 3}, {4}},
		},
		{
			name: "1-5 sum to 15",
			args: args{
				numbers: []int{1, 2, 3, 4, 5},
				k:       15,
			},
			want: [][]int{{1, 2, 3, 4, 5}},
		},
		{
			name: "1-10 sum to 10",
			args: args{
				numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				k:       10,
			},
			want: [][]int{
				{1, 2, 3, 4}, {1, 2, 7}, {1, 3, 6},
				{1, 4, 5}, {1, 9}, {2, 3, 5},
				{2, 8}, {3, 7}, {4, 6}, {10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumToK(tt.args.numbers, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumToK() = %v, want %v", got, tt.want)
			}
		})
	}
}
