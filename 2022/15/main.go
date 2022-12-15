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

func (c *Coord) Rotate90Clockwise() {
	c.x, c.y = c.y, -c.x
}

func (c *Coord) Add(o *Coord) Coord {
	return Coord{x: c.x + o.x, y: c.y + o.y}
}

type Circle struct {
	center Coord
	radius int
}

func (c *Circle) contains(coord *Coord) bool {
	return c.center.Manhattan(coord) <= c.radius
}

func (c *Circle) traceEdge(min, max int) []Coord {
	// returns an array of coords starting from the top and working clockwise
	// provided the x,y coords are in the range min,max
	out := make([]Coord, 0)
	expandedR := c.radius + 1
	maxSteps := 4 * expandedR
	initialPos := Coord{c.center.x + expandedR, c.center.y}
	dir := Coord{1, -1}
	for i := 0; i < maxSteps; i++ {
		if min <= initialPos.x && initialPos.x <= max && min <= initialPos.y && initialPos.y <= max {
			out = append(out, initialPos)
		}
		if i%expandedR == 0 {
			dir.Rotate90Clockwise()
		}
		initialPos = initialPos.Add(&dir)
	}
	return out
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
	fmt.Println("Part 2: ", Part2(data, 4000000))
}

type Interval struct {
	start    int
	end      int
	negative bool
}

func (i *Interval) intersection(o *Interval) *Interval {
	a, b := i.start, i.end
	c, d := o.start, o.end

	// make a <= c
	if c < a {
		tempA := a
		a = c
		c = tempA
		tempB := b
		b = d
		d = tempB
	}

	if b < c {
		return nil
	} else if b < d {
		return &Interval{start: c, end: b, negative: !o.negative}
	} else {
		return &Interval{start: c, end: d, negative: !o.negative}
	}
}

func Part1(sensors []Sensor, targetRow int) int {
	// should update this to store line intervals instead
	// then we can just combine and split intervals
	// or just indicate +- and reduce at the end
	occupiedInTargetRow := make([]Interval, 0)
	beaconsInRow := map[int]bool{}
	for _, sensor := range sensors {
		if sensor.closestBeacon.y == targetRow {
			beaconsInRow[sensor.closestBeacon.x] = true
		}

		radius := sensor.position.Manhattan(&sensor.closestBeacon)

		gap := util.IntAbs(sensor.position.y - targetRow)

		if gap > radius {
			continue
		}

		xFlex := radius - gap

		newInterval := Interval{start: sensor.position.x - xFlex, end: sensor.position.x + xFlex, negative: false}

		toAdd := make([]Interval, 1)
		toAdd[0] = newInterval

		for _, interval := range occupiedInTargetRow {
			intersection := newInterval.intersection(&interval)
			if intersection != nil {
				toAdd = append(toAdd, *intersection)
			}
		}
		occupiedInTargetRow = append(occupiedInTargetRow, toAdd...)
	}

	totalOccupied := 0
	for _, interval := range occupiedInTargetRow {
		distance := interval.end - interval.start + 1
		if interval.negative {
			distance = -distance
		}
		totalOccupied += distance
	}
	totalOccupied -= len(beaconsInRow)

	return totalOccupied
}

func Part2(sensors []Sensor, max int) int {
	circles := make([]*Circle, len(sensors))

	for i, sensor := range sensors {
		circles[i] = &Circle{center: sensor.position, radius: sensor.position.Manhattan(&sensor.closestBeacon)}
	}

	for _, circle := range circles {
		edge := circle.traceEdge(0, max)
		for _, pos := range edge {
			if isContained(&pos, &circles) {
				continue
			}
			return pos.x*4000000 + pos.y
		}
	}

	return -1
}

func isContained(pos *Coord, circles *[]*Circle) bool {
	for _, circle := range *circles {
		if circle.contains(pos) {
			return true
		}
	}
	return false

}
