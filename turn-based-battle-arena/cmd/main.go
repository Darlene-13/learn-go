package main

import (
	"fmt"

	"github.com/Darlene-13/learn-go/turn-based-battle-arena/internal/battle"
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

	warrior.Attack(&mage) // Making them fight
	fmt.Println("Mage HP:", mage.HP)

	healer.Attack(&warrior)
	fmt.Println("Warrior HP MAX:", warrior.MaxHP)
	fmt.Println("Warrior HP:", warrior.HP)

	healer.Attack(&warrior)
	fmt.Println("Warrior HP MAX:", warrior.MaxHP)
	fmt.Println("Warrior HP:", warrior.HP)

	mage.Attack(&warrior)
	fmt.Println("Warrior HP MAX:", warrior.MaxHP)
	fmt.Println("Warrior HP:", warrior.HP)

	fighters := []fighter.Fighter{
		&healer,
		&mage,
		&warrior,
	}

	teamA := []fighter.Fighter{
		&warrior,
	}

	teamB := []fighter.Fighter{
		&mage,
	}

	b := battle.Battle{
		TeamA: teamA,
		TeamB: teamB,
	}

	b.RunRound()

	fmt.Println("Warrior alive: ", warrior.IsAlive())
	fmt.Println("Mage alive: ", mage.IsAlive())
	fmt.Println("Mage HP: ", mage.HP)

	for _, f := range fighters {
		fmt.Println(f.IsAlive())
	}
}
