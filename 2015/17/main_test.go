package main

import "testing"

func Test_dp2(t *testing.T) {
	type args struct {
		file      string
		targetNum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test1", args: args{file: "test.txt", targetNum: 25}, want: 3},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dp2(tt.args.file, tt.args.targetNum); got != tt.want {
				t.Errorf("dp2() = %v, want %v", got, tt.want)
			}
		})
	}
}
