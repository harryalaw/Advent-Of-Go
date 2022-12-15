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

// returns an array of coords starting from the top and working clockwise
// provided the x,y coords are in the range min,max
func (c *Circle) traceEdge(min, max int) []Coord {
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

// Line of the form y = mx + c
// if we know y0 and y1 we can compute x0,x1 if needed
// probably use x0,x1
// where m is ±1
type Line struct {
	x0 int
	x1 int
	m  int
	c  int
}

// l = ax + c
// o = bx + d
// a == 1 as l is posLine
// b == -1 as o is a negLine
func (l *Line) intersect(o *Line) *Coord {
	x := (o.c - l.c)
	if x%2 != 0 {
		return nil
	}
	x /= 2
	if l.x0 <= x && x <= l.x1 && o.x0 <= x && x <= o.x1 {
		return &Coord{x: x, y: x + l.c}
	}
	return nil
}

// there are four lines that bound the circle
// (x-r, y) -> (x, y+r)
// (x-r, y) -> (x, y-r)
// (x, y+r) -> (x+r, y)
// (x, y-r) -> (x+r, y)
// these all have slope ±1
// and c is given by initial x,y values as c = y0 - mx0 <=> c= y0 ± x0
func (c *Circle) getLines() (Line, Line, Line, Line) {
	x := c.center.x
	y := c.center.y
	r := c.radius

	// (x-r, y) -> (x, y+r)
	l1 := Line{x0: x - r, x1: x, m: 1, c: y - x + r}
	// (x, y-r) -> (x+r, y)
	l2 := Line{x0: x, x1: x + 1, m: 1, c: y - x - r}
	// (x, y+r) -> (x+r, y)
	l3 := Line{x0: x, x1: x + r, m: -1, c: y + x - r}
	// (x, y-r) -> (x+r, y)
	l4 := Line{x0: x, x1: x + r, m: -1, c: y + x + r}

	return l1, l2, l3, l4
}

func Part2(sensors []Sensor, max int) int {
	circles := make([]*Circle, len(sensors))
	// how can we exploit our knowledge for what I'm solving
	// there is only one item that is not in a circle
	// so it must be bounded on two sides by a line
	// there must be two such lines that have this property
	// then we can compute the intersection of the two lines
	// Instead of getting all the edges we could instead take a circle and break it down into 4 lines
	// then its enough to find all pairs of lines which have a 1 unit gap between them
	// then there's probably only one good item?

	positiveGradientLines := map[int][]Line{}
	negativeGradientLines := map[int][]Line{}

	for i, sensor := range sensors {
		circles[i] = &Circle{center: sensor.position, radius: sensor.position.Manhattan(&sensor.closestBeacon)}
		l1, l2, l3, l4 := circles[i].getLines()

		positiveGradientLines[l1.c] = append(positiveGradientLines[l1.c], l1)
		positiveGradientLines[l2.c] = append(positiveGradientLines[l2.c], l2)

		negativeGradientLines[l3.c] = append(negativeGradientLines[l3.c], l3)
		negativeGradientLines[l4.c] = append(negativeGradientLines[l4.c], l4)
	}

	// for each positive line we need to check if there is a line with offset ±2
	possiblePosLines := make([]Line, 0)

	for offset := range positiveGradientLines {
		_, ok := positiveGradientLines[offset+2]
		if !ok {
			continue
		}
		possiblePosLines = append(possiblePosLines, Line{
			x0: 0,
			x1: max,
			m:  1,
			c:  offset + 1,
		})
	}

	possibleNegLines := make([]Line, 0)

	for offset := range negativeGradientLines {
		possibleNegLines = append(possibleNegLines, Line{
			x0: 0,
			x1: max,
			m:  -1,
			c:  offset + 1,
		})
	}

	intersectionPoints := make([]*Coord, 0)

	for _, posLine := range possiblePosLines {
		for _, negLine := range possibleNegLines {
			intersectionPoint := posLine.intersect(&negLine)
			if intersectionPoint != nil {
				intersectionPoints = append(intersectionPoints, intersectionPoint)
			}
		}
	}

	for _, point := range intersectionPoints {
		if isContained(point, &circles) {
			continue
		}
		return point.x*4000000 + point.y
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
