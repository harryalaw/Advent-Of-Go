package main

import (
	"reflect"
	"testing"
)

func Test_getNeighbours(t *testing.T) {
	type args struct {
		point  coord
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want []coord
	}{
		{
			"0 0",
			args{point: coord{x: 0, y: 0}, height: 100, width: 100},
			[]coord{{x: 0, y: 1}, {x: 1, y: 0}, {x: 1, y: 1}},
		},
		{
			"5 5",
			args{point: coord{x: 5, y: 5}, height: 100, width: 100},
			[]coord{{x: 4, y: 4}, {x: 4, y: 5}, {x: 4, y: 6}, {x: 5, y: 4}, {x: 5, y: 6}, {x: 6, y: 4}, {x: 6, y: 5}, {x: 6, y: 6}},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNeighbours(tt.args.point, tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNeighbours() = %v, want %v", got, tt.want)
			}
		})
	}
}
