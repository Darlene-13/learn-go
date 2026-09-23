package fighter

import "strconv"

type Warrior struct {
	HP          int
	MaxHP       int
	AttackPower int
	Name        string
}

func (w Warrior) Attack(target Fighter) string {
	if !target.IsAlive() {
		return "Cannot attack: target is already dead"
	}
	target.TakeDamage(w.AttackPower)

	// strconv.Itoa go method to convert int to integer
	return w.Name + "Attacks for: " + strconv.Itoa(w.AttackPower) + "damage"

}

// TakeDamage We are using a pointer because it is going to modify the actual warrior and not a copy of it.
func (w *Warrior) TakeDamage(damage int) int {
	w.HP -= damage

	if w.HP < 0 {
		w.HP = 0
	}

	return w.HP
}

func (w Warrior) IsAlive() bool {
	return w.HP > 0
}
