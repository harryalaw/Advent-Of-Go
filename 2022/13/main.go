package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"sort"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

type Result int

const (
	Greater Result = 1
	Equal   Result = 0
	Less    Result = -1
)

type Value interface {
	LessThan(v Value) Result
	String() string
}

type Packet struct {
	data Value
}

type List struct {
	items []Value
}

func (l *List) String() string {
	var out bytes.Buffer
	out.WriteRune('[')
	for i, item := range l.items {
		out.WriteString(item.String())
		if i != len(l.items)-1 {
			out.WriteRune(',')
		}
	}
	out.WriteRune(']')

	return out.String()
}

func (l *List) LessThan(v Value) Result {
	switch v := v.(type) {
	case *List:
		minLength := util.IntMin(len(l.items), len(v.items))
		for i := 0; i < minLength; i++ {
			res := l.items[i].LessThan(v.items[i])
			if res == Less || res == Greater {
				return res
			}
		}
		if len(l.items) < len(v.items) {
			return Less
		} else if len(l.items) > len(v.items) {
			return Greater
		}
		return Equal
	case *Integer:
		newList := v.ToList()
		return l.LessThan(&newList)
	default:
		panic("Not implemented")
	}
}

type Integer int

func (i *Integer) String() string {
	return fmt.Sprintf("%d", i)
}

func (i *Integer) ToList() List {
	newList := make([]Value, 0)
	newList = append(newList, i)

	return List{items: newList}
}

func (i *Integer) LessThan(v Value) Result {
	switch v := v.(type) {
	case *List:
		newList := i.ToList()
		return newList.LessThan(v)
	case *Integer:
		if *i == *v {
			return Equal
		} else if *i < *v {
			return Less
		}
		return Greater
	default:
		panic("Not implemented")
	}
}

type Packets struct {
	left  Packet
	right Packet
}

func isDigit(r byte) bool {
	return r == '0' ||
		r == '1' ||
		r == '2' ||
		r == '3' ||
		r == '4' ||
		r == '5' ||
		r == '6' ||
		r == '7' ||
		r == '8' ||
		r == '9'
}
func parseInt(text string) (Integer, int) {
	i := 0
	for isDigit(text[i]) {
		i++
	}
	number := Integer(util.Atoi(text[0:i]))

	return number, i
}

func parseList(text string, list *List) int {
	for i := 0; i < len(text); i++ {
		c := text[i]
		switch c {
		case ',':
			continue
		case ']':
			return i + 1
		case '[':
			newList := List{items: []Value{}}
			offset := parseList(text[i+1:], &newList)
			list.items = append(list.items, &newList)
			i += offset
		default:
			number, offset := parseInt(text[i:])
			i += offset - 1
			list.items = append(list.items, &number)
		}
	}
	return len(text)
}

func value(text string) Value {
	c := text[0]
	if c == '[' {
		newList := List{items: []Value{}}
		parseList(text[1:], &newList)
		return &newList
	} else {
		number, _ := parseInt(text)
		return &number
	}
}

func parseInput(input string) []Packets {
	data := strings.Split(strings.TrimSpace(input), "\n\n")

	out := make([]Packets, len(data))

	for i, pair := range data {
		lines := strings.Split(pair, "\n")
		leftText := lines[0]
		rightText := lines[1]

		leftData := value(leftText)
		rightData := value(rightText)

		out[i] = Packets{
			left:  Packet{data: leftData},
			right: Packet{data: rightData},
		}
	}

	return out
}

func main() {
	util.Time(doPart1, "Part 1")
	util.Time(doPart2, "Part 2")
}

func doPart1() {
	data := parseInput(input)
	fmt.Println("Part 1: ", Part1(data))
}

func doPart2() {
	data := parseInput(input)
	fmt.Println("Part 2: ", Part2(data))
}

func Part1(packets []Packets) int {
	total := 0
	for i, packet := range packets {
		result := packet.left.data.LessThan(packet.right.data)
		if result == Less {
			total += i + 1
		}
	}
	return total
}

func Part2(data []Packets) int {
	packets := make([]Value, len(data)*2)
	for i, value := range data {
		packets[2*i] = value.left.data
		packets[2*i+1] = value.right.data
	}

	divider2 := value("[[2]]")
	divider6 := value("[[6]]")

	packets = append(packets, divider2, divider6)

	sort.Slice(packets, func(i, j int) bool {
		return packets[i].LessThan(packets[j]) == Less
	})

	total := 1
	for i := 0; i < len(packets); i++ {
		if packets[i].String() == divider2.String() || packets[i].String() == divider6.String() {
			total *= (i + 1)
		}
	}
	return total
}
