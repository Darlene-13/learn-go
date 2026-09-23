package fighter

import "strconv"

type Mage struct {
	HP          int
	MaxHP       int
	AttackPower int
	Name        string
}

func (m Mage) Attack(target Fighter) string {
	// Check if they are alive
	if !m.IsAlive() {
		return "Cannot attack: target is already dead"
	}

	//Calculate the damage it is supposed to cause
	damage := m.AttackPower * 2

	//Let the target be damaged by the damage rate
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
