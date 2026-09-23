package fighter

// Fighter defines the behavior every fighter must provide.

type Fighter interface {
	Attack(target Fighter) string
	TakeDamage(damage int) int
	IsAlive() bool
}
