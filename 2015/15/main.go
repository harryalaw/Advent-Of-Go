package main

import (
	util "2015/Util"
	"fmt"
)

func main() {
	util.Time(part1, "Part1")
	util.Time(part2, "Part2")
}
func part1() {
	fmt.Println(brute())
}
func part2() {
	fmt.Println(brute2())
}

func brute() int {
	maxPoints := -1
	for x := 0; x <= 100; x++ {
		for y := 0; y <= 100-x; y++ {
			for z := 0; z <= 100-x-y; z++ {
				w := 100 - x - y - z
				temp := f(x, y, z, w)
				if temp > maxPoints {
					maxPoints = temp
				}
			}
		}
	}
	return maxPoints
}

func brute2() int {
	maxPoints := -1
	for x := 0; x <= 100; x++ {
		for y := 0; y <= 100-x; y++ {
			for z := 0; z <= 100-x-y; z++ {
				w := 100 - x - y - z
				if cal(x, y, z, w) != 500 {
					continue
				}
				temp := f(x, y, z, w)
				if temp > maxPoints {
					maxPoints = temp
				}
			}
		}
	}
	return maxPoints
}

func cal(x, y, z, w int) int {
	return 5*x + y + 6*z + 8*w
}

func f(x, y, z, w int) int {
	if 5*x-y-w < 0 {
		return 0
	}
	if 3*y-x-z < 0 {
		return 0
	}
	return (5*x - y - w) * (3*y - x - z) * 4 * z * 2 * w
}
