package main

import (
	"fmt"

	"github.com/Darlene-13/learn-go/turn-based-battle-arena/internal/fighter"
)

func main() {

	warrior := fighter.Warrior{
		Name:        "Darlene",
		HP:          100,
		MaxHP:       100,
		AttackPower: 20,
	}

	mage := fighter.Mage{
		Name:        "Wendy",
		HP:          80,
		MaxHP:       80,
		AttackPower: 15,
	}

	healer := fighter.Healer{
		Name:        "Stacy",
		HP:          70,
		MaxHP:       70,
		AttackPower: 10,
	}

	fighters := []fighter.Fighter{
		&healer,
		&mage,
		&warrior,
	}

	for _, f := range fighters {
		fmt.Println(f.IsAlive())
	}
}
