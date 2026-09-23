package battle

import "github.com/Darlene-13/learn-go/turn-based-battle-arena/internal/fighter"

type Battle struct {
	TeamA []fighter.Fighter
	TeamB []fighter.Fighter
}

func (b Battle) IsTeamAlive(team []fighter.Fighter) bool {
	for _, f := range team {
		if f.IsAlive() {
			return true
		}
	}
	return false
}

func (b Battle) FindTarget(team []fighter.Fighter) (fighter.Fighter, bool) {
	for _, f := range team {
		if f.IsAlive() {
			return f, true
		}
	}

	return nil, false
}

func (b Battle) RunRound() {
	// Team A attacks
	for _, team := range b.TeamA {
		if !team.IsAlive() {
			continue
		}

		target, found := b.FindTarget(b.TeamB)

		if !found {
			break
		}

		team.Attack(target)
	}

	// Team B attacks
	for _, team := range b.TeamB {
		if !team.IsAlive() {
			continue
		}

		target, found := b.FindTarget(b.TeamA)

		if !found {
			break
		}

		team.Attack(target)
	}
}
