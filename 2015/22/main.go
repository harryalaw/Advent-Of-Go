package main

import (
	"github.com/harryalaw/advent-of-go/util"
	"fmt"
	"io/ioutil"
	"strings"
)

var allSpells []Spell

const MAX_INT = 0xffffffff

func main() {
	allSpells = []Spell{
		{name: "MagicMissile", manaCost: 53},
		{name: "Drain", manaCost: 73},
		{name: "Shield", manaCost: 113},
		{name: "Poison", manaCost: 173},
		{name: "Recharge", manaCost: 229},
	}
	util.Time(doPart1, "Part1")
	util.Time(doPart2, "Part2")
}

type Player struct {
	hp        int
	mana      int
	armour    int
	manaSpent int
}

type Boss struct {
	hp     int
	damage int
}

type Spell struct {
	name              string
	manaCost          int
	remainingDuration int
}

func parseInput(file string) Boss {
	content, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	return Boss{
		hp:     util.Atoi(strings.Fields(lines[0])[2]),
		damage: util.Atoi(strings.Fields(lines[1])[1]),
	}
}

func doPart1() {
	boss := parseInput("data.txt")
	player := Player{hp: 50, mana: 500, manaSpent: 0}

	fmt.Println(part1(player, boss))
}

func doPart2() {
	boss := parseInput("data.txt")
	player := Player{hp: 50, mana: 500, manaSpent: 0}

	fmt.Println(part2(player, boss))
}

func part1(player Player, boss Boss) int {
	return playGame(player, boss, false)
}

func part2(player Player, boss Boss) int {
	return playGame(player, boss, true)
}

func playGame(player Player, boss Boss, hardMode bool) int {
	allSpells = []Spell{
		{name: "MagicMissile", manaCost: 53},
		{name: "Drain", manaCost: 73},
		{name: "Shield", manaCost: 113},
		{name: "Poison", manaCost: 173},
		{name: "Recharge", manaCost: 229},
	}
	activeSpells := make(map[string]Spell)

	minCost := MAX_INT
	doTurn(player, boss, activeSpells, true, &minCost, hardMode)
	return minCost
}

func applyEffect(spell *Spell, player Player, boss Boss) (Player, Boss) {
	if spell.remainingDuration <= 0 {
		return player, boss
	}
	spell.remainingDuration -= 1
	switch spell.name {
	case "Shield":
		player.armour = 7
	case "Poison":
		boss.hp -= 3
	case "Recharge":
		player.mana += 101
	}

	return player, boss
}

func doTurn(player Player,
	boss Boss,
	activeSpells map[string]Spell,
	isPlayerTurn bool,
	minCost *int,
	hardMode bool) int {

	if hardMode && isPlayerTurn {
		player.hp -= 1
		if player.hp <= 0 {
			return MAX_INT
		}
	}

	player.armour = 0
	for _, spell := range activeSpells {
		player, boss = applyEffect(&spell, player, boss)

		activeSpells[spell.name] = spell
	}

	if boss.hp <= 0 {
		if player.manaSpent < *minCost {
			*minCost = player.manaSpent
		}
		return player.manaSpent
	}

	if isPlayerTurn {
		return doPlayerTurn(player, boss, activeSpells, minCost, hardMode)
	} else {
		return doBossTurn(player, boss, activeSpells, minCost, hardMode)
	}
}

func doBossTurn(player Player,
	boss Boss,
	activeSpells map[string]Spell,
	minCost *int,
	hardMode bool) int {
	damage := boss.damage - player.armour
	if damage < 1 {
		damage = 1
	}
	player.hp -= damage

	if player.hp <= 0 {
		return MAX_INT
	}
	return doTurn(player, boss, activeSpells, true, minCost, hardMode)

}

func doPlayerTurn(player Player,
	boss Boss,
	activeSpells map[string]Spell,
	minCost *int,
	hardMode bool) int {
	if player.mana < 53 {
		return MAX_INT
	}

	tempMin := MAX_INT
	for _, spell := range allSpells {
		tempPlayer := player
		if tempPlayer.mana >= spell.manaCost && notActive(spell, activeSpells) {
			tempPlayer, tempBoss, tempActiveSpells := castSpell(spell, player, boss, activeSpells)
			if tempPlayer.manaSpent > *minCost {
				continue
			}
			tempCost := doTurn(tempPlayer, tempBoss, tempActiveSpells, false, minCost, hardMode)

			if tempCost < tempMin {
				tempMin = tempCost
			}
		}
	}
	return tempMin
}

func castSpell(spell Spell, player Player, boss Boss, activeSpells map[string]Spell) (Player, Boss, map[string]Spell) {
	newActiveSpells := make(map[string]Spell)

	for key, value := range activeSpells {
		newActiveSpells[key] = value
	}

	player.mana -= spell.manaCost
	player.manaSpent += spell.manaCost

	switch spell.name {
	case "MagicMissile":
		boss.hp -= 4
	case "Drain":
		player.hp += 2
		boss.hp -= 2
	case "Shield":
		newActiveSpells["Shield"] = Spell{
			name:              "Shield",
			manaCost:          113,
			remainingDuration: 6,
		}
	case "Poison":
		newActiveSpells["Poison"] = Spell{
			name:              "Poison",
			manaCost:          173,
			remainingDuration: 6,
		}
	case "Recharge":
		newActiveSpells["Recharge"] = Spell{
			name:              "Recharge",
			manaCost:          101,
			remainingDuration: 5,
		}
	}

	return player, boss, newActiveSpells
}

func notActive(spell Spell, activeSpells map[string]Spell) bool {
	spellInfo, isPresent := activeSpells[spell.name]

	return !isPresent || spellInfo.remainingDuration == 0
}
