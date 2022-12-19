package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
	"time"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

func parseInput(input string) []Blueprint {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	blueprints := make([]Blueprint, len(lines))

	for i, line := range lines {
		robotCosts := strings.Split(strings.Split(line, ": ")[1], ". ")

		oreCost := util.Atoi(strings.Fields(robotCosts[0])[4])

		clayCost := util.Atoi(strings.Fields(robotCosts[1])[4])

		obsidianCosts := strings.Fields(robotCosts[2])
		obsidianOreCost := util.Atoi(obsidianCosts[4])
		obsidianClayCost := util.Atoi(obsidianCosts[7])

		geodeCosts := strings.Fields(robotCosts[3])
		geodeOreCost := util.Atoi(geodeCosts[4])
		geodeObsidianCost := util.Atoi(geodeCosts[7])

		blueprints[i] = Blueprint{
			oreCost:      Cost{oreCost, 0, 0},
			clayCost:     Cost{clayCost, 0, 0},
			obsidianCost: Cost{obsidianOreCost, obsidianClayCost, 0},
			geodeCost:    Cost{geodeOreCost, 0, geodeObsidianCost},
		}
	}

	return blueprints
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

type Cost struct {
	ore      int
	clay     int
	obsidian int
}

type Blueprint struct {
	oreCost      Cost
	clayCost     Cost
	obsidianCost Cost
	geodeCost    Cost
}

func Part1(blueprints []Blueprint) int {
	qualityTotal := 0
	for i, blueprint := range blueprints {
		qualityLevel := ProbabilisticBlueprint(blueprint, 10000, 24)
		qualityLevel *= (i + 1)
		qualityTotal += qualityLevel
	}
	return qualityTotal
}

func Part2(blueprints []Blueprint) int {
	total := 1
	for i := 0; i < 3; i++ {
		blueprint := blueprints[i]
		qualityLevel := ProbabilisticBlueprint(blueprint, 100000, 32)
		// Running it several times to confirm it works
		fmt.Println(i, qualityLevel)
		total *= qualityLevel
	}
	return total
}

type State struct {
	ore           int
	clay          int
	obsidian      int
	geode         int
	robotOre      int
	robotClay     int
	robotObsidian int
	robotGeode    int
}

// should be able to do it with a DFS
func CalculateBlueprint(blueprint Blueprint) int {
	states := map[State]struct{}{}
	initialState := State{robotOre: 1}
	states[initialState] = struct{}{}

	maxClay := blueprint.obsidianCost.clay
	maxOre := util.IntMax(blueprint.oreCost.ore, blueprint.clayCost.ore, blueprint.obsidianCost.ore, blueprint.geodeCost.ore)

	for i := 1; i <= 24; i++ {
		nextStates := map[State]struct{}{}
		for state := range states {
			newStates := spendGoods(state, blueprint, maxOre, maxClay)
			for _, newState := range newStates {
				newState = updateGoods(newState, state)
				nextStates[newState] = struct{}{}
			}
		}
		states = nextStates
	}

	maxGeodes := 0
	for state := range states {
		geodes := state.geode
		if geodes > maxGeodes {
			maxGeodes = geodes
		}
	}

	return maxGeodes
}

func updateGoods(newState State, prevState State) State {
	newState.ore += prevState.robotOre
	newState.clay += prevState.robotClay
	newState.obsidian += prevState.robotObsidian
	newState.geode += prevState.robotGeode

	return newState
}

func spendGoods(state State, blueprint Blueprint, maxOreRobots, maxClayRobots int) []State {
	if state.obsidian >= blueprint.geodeCost.obsidian && state.ore >= blueprint.geodeCost.ore {
		newState := State{
			ore:           state.ore - blueprint.geodeCost.ore,
			clay:          state.clay,
			obsidian:      state.obsidian - blueprint.geodeCost.obsidian,
			geode:         state.geode,
			robotOre:      state.robotOre,
			robotClay:     state.robotClay,
			robotObsidian: state.robotObsidian,
			robotGeode:    state.robotGeode + 1,
		}
		return []State{newState}
	}

	if state.clay >= blueprint.obsidianCost.clay && state.ore >= blueprint.obsidianCost.ore {
		newState := State{
			ore:           state.ore - blueprint.obsidianCost.ore,
			clay:          state.clay - blueprint.obsidianCost.clay,
			obsidian:      state.obsidian,
			geode:         state.geode,
			robotOre:      state.robotOre,
			robotClay:     state.robotClay,
			robotObsidian: state.robotObsidian + 1,
			robotGeode:    state.robotGeode,
		}
		return []State{newState}
	}

	out := make([]State, 0)
	// here is an assumption for pruning
	if state.ore < 10 && state.clay < 10 {
		out = append(out, state)
	}

	if state.robotOre < maxOreRobots && state.ore >= blueprint.oreCost.ore {
		newState := State{
			ore:           state.ore - blueprint.oreCost.ore,
			clay:          state.clay,
			obsidian:      state.obsidian,
			geode:         state.geode,
			robotOre:      state.robotOre + 1,
			robotClay:     state.robotClay,
			robotObsidian: state.robotObsidian,
			robotGeode:    state.robotGeode,
		}
		out = append(out, newState)
	}
	if state.robotClay < maxClayRobots && state.ore >= blueprint.clayCost.ore {
		newState := State{
			ore:           state.ore - blueprint.clayCost.ore,
			clay:          state.clay,
			obsidian:      state.obsidian,
			geode:         state.geode,
			robotOre:      state.robotOre,
			robotClay:     state.robotClay + 1,
			robotObsidian: state.robotObsidian,
			robotGeode:    state.robotGeode,
		}
		out = append(out, newState)
	}

	return out
}

func ProbabilisticBlueprint(blueprint Blueprint, iters int, timer int) int {
	maxClay := blueprint.obsidianCost.clay
	maxOre := util.IntMax(blueprint.oreCost.ore, blueprint.clayCost.ore, blueprint.obsidianCost.ore, blueprint.geodeCost.ore)

	maxGeodes := 0
	for iter := 0; iter < iters; iter++ {
		state := simulateBlueprint(blueprint, maxOre, maxClay, timer)
		if state.geode > maxGeodes {
			maxGeodes = state.geode
		}
	}

	return maxGeodes
}

func simulateBlueprint(blueprint Blueprint, maxOre, maxClay int, timer int) State {
	state := State{robotOre: 1}
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= timer; i++ {
		newState := probableSpend(state, blueprint, maxClay, maxOre)
		newState = updateGoods(newState, state)
		state = newState
	}

	return state
}

func probableSpend(state State, blueprint Blueprint, maxOre int, maxClay int) State {
	if state.obsidian >= blueprint.geodeCost.obsidian && state.ore >= blueprint.geodeCost.ore {
		return State{
			ore:           state.ore - blueprint.geodeCost.ore,
			clay:          state.clay,
			obsidian:      state.obsidian - blueprint.geodeCost.obsidian,
			geode:         state.geode,
			robotOre:      state.robotOre,
			robotClay:     state.robotClay,
			robotObsidian: state.robotObsidian,
			robotGeode:    state.robotGeode + 1,
		}
	} else if state.clay >= blueprint.obsidianCost.clay && state.ore >= blueprint.obsidianCost.ore && rand.Float32() < float32(0.8) {
		return State{
			ore:           state.ore - blueprint.obsidianCost.ore,
			clay:          state.clay - blueprint.obsidianCost.clay,
			obsidian:      state.obsidian,
			geode:         state.geode,
			robotOre:      state.robotOre,
			robotClay:     state.robotClay,
			robotObsidian: state.robotObsidian + 1,
			robotGeode:    state.robotGeode,
		}
	}

	if state.ore >= blueprint.clayCost.ore && rand.Float32() < float32(0.5) {
		return State{
			ore:           state.ore - blueprint.clayCost.ore,
			clay:          state.clay,
			obsidian:      state.obsidian,
			geode:         state.geode,
			robotOre:      state.robotOre,
			robotClay:     state.robotClay + 1,
			robotObsidian: state.robotObsidian,
			robotGeode:    state.robotGeode,
		}
	}

	if state.ore >= blueprint.oreCost.ore && rand.Float32() < float32(0.5) {
		return State{
			ore:           state.ore - blueprint.oreCost.ore,
			clay:          state.clay,
			obsidian:      state.obsidian,
			geode:         state.geode,
			robotOre:      state.robotOre + 1,
			robotClay:     state.robotClay,
			robotObsidian: state.robotObsidian,
			robotGeode:    state.robotGeode,
		}
	}

	return state
}
