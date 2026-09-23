package fighter

import "strconv"

type Mage struct {
	HP          int
	MaxHP       int
	AttackPower int
	Name        string
}

func (m Mage) Attack(target Fighter) string {
	if !m.IsAlive() {
		return "Cannot attack: target is already dead"
	}

	damage := m.AttackPower * 2

	target.TakeDamage(damage)

	return m.Name + "Attacks for " + strconv.Itoa(damage) + "damage"

}

func (m *Mage) TakeDamage(damage int) int {
	m.HP -= damage

	if m.HP < 0 {
		m.HP = 0
	}

	return m.HP
}

func (m Mage) IsAlive() bool {
	return m.HP > 0
}
