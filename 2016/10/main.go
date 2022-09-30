package main

import (
	_ "embed"
	"fmt"
	"strings"

	util "github.com/harryalaw/advent-of-go/util"
)

//go:embed data.txt
var input string

//go:embed data_test.txt
var test_input string

func main() {
	util.Time(doPart1, "Part1")
	util.Time(doPart2, "")

}

func doPart1() {
	fmt.Println(coreLogic(strings.TrimSpace(input), true))
}
func doPart2() {
	fmt.Println(coreLogic(strings.TrimSpace(input), false))
}

func parseText(input string) map[int]Bot {
	bots := make(map[int]Bot)

	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		switch fields[0] {
		case "bot":
			id := util.Atoi(fields[1])
			lowId := util.Atoi(fields[6])
			highId := util.Atoi(fields[11])

			lowTargetType := fields[5]
			highTargetType := fields[10]

			bot := fetchBot(id, bots)

			bot.lowTarget = getTarget(lowTargetType, lowId)
			bot.highTarget = getTarget(highTargetType, highId)

			bots[id] = bot
		case "value":
			value := util.Atoi(fields[1])
			botId := util.Atoi(fields[5])

			bot := fetchBot(botId, bots)

			bot.addToValues(value)

			bots[botId] = bot
		}
	}
	return bots
}

func fetchBot(id int, bots map[int]Bot) Bot {
	bot, botExists := bots[id]

	if !botExists {
		bot = Bot{id: id}
	}

	return bot
}

func getTarget(typeString string, id int) Target {
	isBot := false
	if typeString == "bot" {
		isBot = true
	}
	return Target{id: id, isBot: isBot}
}

func coreLogic(input string, isPart1 bool) int {
	bots := parseText(input)

	outputs := make(map[int]int)

	changes := true
	for changes {
		changes = false
		for _, bot := range bots {
			if bot.values[0] != 0 && bot.values[1] != 0 {
				changes = true
				lowValue := util.IntMin(bot.values[0], bot.values[1])
				highValue := util.IntMax(bot.values[0], bot.values[1])
				bot.values = [2]int{0, 0}
				bots[bot.id] = bot

				if isPart1 && lowValue == 17 && highValue == 61 {
					return bot.id
				}

				resolveTarget(bot.lowTarget, lowValue, bots, outputs)
				resolveTarget(bot.highTarget, highValue, bots, outputs)
			}
		}
	}

	return outputs[0] * outputs[1] * outputs[2]
}

func resolveTarget(target Target, value int, bots map[int]Bot, outputs map[int]int) (map[int]Bot, map[int]int) {
	if target.isBot {
		targetBot := bots[target.id]
		targetBot.addToValues(value)
		bots[target.id] = targetBot
	} else {
		outputs[target.id] = value
	}
	return bots, outputs
}

type Target struct {
	id    int
	isBot bool
}

type Bot struct {
	id         int
	values     [2]int
	lowTarget  Target
	highTarget Target
}

func (bot *Bot) addToValues(value int) {
	if bot.values[0] == 0 {
		bot.values[0] = value
	} else {
		bot.values[1] = value
	}
}
