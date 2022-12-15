package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

type Coord struct {
	x int
	y int
}

func (c *Coord) Manhattan(o *Coord) int {
	return util.IntAbs(c.x-o.x) + util.IntAbs(c.y-o.y)
}

type Sensor struct {
	position      Coord
	closestBeacon Coord
}

func parseInput(input string) []Sensor {
	data := strings.Split(strings.TrimSpace(input), "\n")
	sensors := make([]Sensor, len(data))

	for i, line := range data {
		re := regexp.MustCompile("(-?\\d+)")

		values := re.FindAllString(line, -1)

		if len(values) != 4 {
			fmt.Println(values)
			panic("Wrong values!")
		}
		sensorPos := Coord{x: util.Atoi(values[0]), y: util.Atoi(values[1])}
		beaconPos := Coord{x: util.Atoi(values[2]), y: util.Atoi(values[3])}

		sensors[i] = Sensor{position: sensorPos, closestBeacon: beaconPos}
	}

	return sensors
}

func main() {
	util.Time(doPart1, "Part 1")
	util.Time(doPart2, "Part 2")
}

func doPart1() {
	data := parseInput(input)
	fmt.Println("Part 1: ", Part1(data, 2000000))
}

func doPart2() {
	data := parseInput(input)
	fmt.Println("Part 2: ", Part2(data))
}

func Part1(sensors []Sensor, targetRow int) int {
	// should update this to store line intervals instead
	// then we can just combine and split intervals
	// or just indicate +- and reduce at the end
	occupiedInTargetRow := map[int]bool{}
	beaconsInRow := map[int]bool{}
	for _, sensor := range sensors {
		if sensor.position.y == targetRow {
			occupiedInTargetRow[sensor.position.x] = true
		}
		if sensor.closestBeacon.y == targetRow {
			occupiedInTargetRow[sensor.closestBeacon.x] = true
			beaconsInRow[sensor.closestBeacon.x] = true
		}

		radius := sensor.position.Manhattan(&sensor.closestBeacon)

		gap := util.IntAbs(sensor.position.y - targetRow)

		if gap > radius {
			continue
		}

		xFlex := radius - gap

		// this should be rethought with intervals instead
		// we can add intervals to a list with intersections and a sign
		// then use inclusion/exclusion to count them better
		for i := -xFlex; i <= xFlex; i++ {
			occupiedInTargetRow[sensor.position.x+i] = true
		}
	}

	return len(occupiedInTargetRow) - len(beaconsInRow)
}

func Part2(sensors []Sensor) int {
	// need to find all the occupied sensors in the box
	// 0 -> 4000000

	distress := Coord{-1, 1}
	return distress.x * distress.y
}
