package main

import (
	"fmt"
	"github.com/harryalaw/advent-of-go/util"
	"io/ioutil"
	"math"
	"strings"
)

type Player struct {
	hp     int
	damage int
	armor  int
}

type Weapon struct {
	cost   int
	damage int
}

type Armor struct {
	cost  int
	armor int
}

type Ring struct {
	cost   int
	damage int
	armor  int
}

func main() {
	util.Time(fight, "Part1 and Part2")
}

func playerWinsFight(player, boss Player) bool {
	bossDamage := boss.damage - player.armor
	if bossDamage <= 1 {
		bossDamage = 1
	}
	roundsTillPlayerDies := 1 + ((player.hp - 1) / bossDamage)

	playerDamage := player.damage - boss.armor
	if playerDamage <= 1 {
		playerDamage = 1
	}
	roundsTillBossDies := 1 + ((boss.hp - 1) / playerDamage)

	return roundsTillBossDies <= roundsTillPlayerDies
}

func populateShop() ([]Weapon, []Armor, []Ring) {
	weapons := make([]Weapon, 5)
	weapons[0] = Weapon{cost: 8, damage: 4}
	weapons[1] = Weapon{cost: 10, damage: 5}
	weapons[2] = Weapon{cost: 25, damage: 6}
	weapons[3] = Weapon{cost: 40, damage: 7}
	weapons[4] = Weapon{cost: 74, damage: 8}

	armors := make([]Armor, 5)
	armors[0] = Armor{cost: 13, armor: 1}
	armors[1] = Armor{cost: 31, armor: 2}
	armors[2] = Armor{cost: 53, armor: 3}
	armors[3] = Armor{cost: 75, armor: 4}
	armors[4] = Armor{cost: 102, armor: 5}

	rings := make([]Ring, 6)
	rings[0] = Ring{cost: 25, damage: 1}
	rings[1] = Ring{cost: 50, damage: 2}
	rings[2] = Ring{cost: 100, damage: 3}
	rings[3] = Ring{cost: 20, armor: 1}
	rings[4] = Ring{cost: 40, armor: 2}
	rings[5] = Ring{cost: 80, armor: 3}

	return weapons, armors, rings
}

func createPlayer(weaponId, armorId, ring1Id, ring2Id int) (Player, int) {
	weapons, armors, rings := populateShop()

	player := Player{hp: 100, damage: 0, armor: 0}
	totalCost := 0

	if ring1Id != -1 && ring1Id == ring2Id {
		return player, totalCost
	}

	if weaponId != -1 {
		player.damage += weapons[weaponId].damage
		totalCost += weapons[weaponId].cost
	}

	if armorId != -1 {
		player.armor += armors[armorId].armor
		totalCost += armors[armorId].cost
	}

	if ring1Id != -1 {
		player.damage += rings[ring1Id].damage
		player.armor += rings[ring1Id].armor
		totalCost += rings[ring1Id].cost
	}

	if ring2Id != -1 {
		player.damage += rings[ring2Id].damage
		player.armor += rings[ring2Id].armor
		totalCost += rings[ring2Id].cost
	}

	return player, totalCost
}

func fight() {
	boss := parseInput()
	minCost := math.MaxInt
	maxCost := -1
	for weaponId := 0; weaponId < 5; weaponId++ {
		for armorId := -1; armorId < 5; armorId++ {
			for ring1Id := -1; ring1Id < 6; ring1Id++ {
				for ring2Id := -1; ring2Id < 6; ring2Id++ {
					player, cost := createPlayer(weaponId, armorId, ring1Id, ring2Id)
					if cost < minCost && playerWinsFight(player, boss) {
						minCost = cost
					}
					if cost > maxCost && !playerWinsFight(player, boss) {
						maxCost = cost
					}
				}
			}
		}
	}
	fmt.Println(minCost)
	fmt.Println(maxCost)
}

func parseInput() Player {
	data, err := ioutil.ReadFile("data.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	return Player{
		hp:     util.Atoi(strings.Fields(lines[0])[2]),
		damage: util.Atoi(strings.Fields(lines[1])[1]),
		armor:  util.Atoi(strings.Fields(lines[2])[1]),
	}

}
