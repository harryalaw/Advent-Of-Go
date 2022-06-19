package main

import "testing"

func Test_doPart1(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{file: "test1.txt"},
			want: 4,
		},
		{
			name: "Test 2",
			args: args{file: "test2.txt"},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doPart1(tt.args.file); got != tt.want {
				t.Errorf("doPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doPart2(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{file: "test1.txt"},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{file: "test2.txt"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doPart2(tt.args.file); got != tt.want {
				t.Errorf("doPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
