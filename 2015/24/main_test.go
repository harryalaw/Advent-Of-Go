package main

import (
	"reflect"
	"testing"
)

func Test_sortDescending(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Descending list stays same",
			args: []int{4, 3, 2, 1},
			want: []int{4, 3, 2, 1},
		},
		{
			name: "Scrambled array becomes sorted",
			args: []int{1, 3, 2, 4},
			want: []int{4, 3, 2, 1},
		},
		{
			name: "Sorted ascending becomes descending",
			args: []int{1, 2, 3, 4},
			want: []int{4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortDescending(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortDescending() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

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

func Test_sumToKPruned(t *testing.T) {
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
			want: [][]int{{4}},
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
			want: [][]int{{10}},
		},
		{
			name: "Longer sum after valid sum",
			args: args{
				numbers: []int{9, 3, 2, 1},
				k:       12,
			},
			want: [][]int{{9, 3}},
		},
		{
			name: "Descending values that don't sum to target",
			args: args{
				numbers: []int{5, 4, 3, 2, 1},
				k:       20,
			},
			want: [][]int{},
		},
		{
			name: "Example from page",
			args: args{
				numbers: []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
				k:       15,
			},
			want: [][]int{{4, 11}, {5, 10}, {7, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumToKPruned(tt.args.numbers, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumToKPruned() = %v, want %v", got, tt.want)
			}
		})
	}
}
