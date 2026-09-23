package fighter

import "strconv"

type Healer struct {
	HP          int
	MaxHP       int
	AttackPower int
	Name        string
}

func (h Healer) Attack(target Fighter) string {
	if !h.IsAlive() {
		return "Cannot attack: target is already dead"
	}

	damage := h.AttackPower / 2
	target.TakeDamage(damage)
	return h.Name + " attacks for " + strconv.Itoa(damage) + " damage"

}

func (h *Healer) TakeDamage(damage int) int {
	h.HP -= damage

	if h.HP < 0 {
		h.HP = 0
	}

	return h.HP
}
func (h Healer) IsAlive() bool {
	return h.HP > 0
}
