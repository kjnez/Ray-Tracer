package tuples

import (
	"fmt"
)

type Projectile struct {
	position Tuple
	velocity Tuple
}

type Environment struct {
	gravity Tuple
	wind    Tuple
}

func tick(env Environment, proj Projectile) Projectile {
	position := Add(proj.position, proj.velocity)
	velocity := Add(Add(proj.velocity, env.gravity), env.wind)
	return Projectile{position, velocity}
}

func SimulateProjectile() {
	p := Projectile{Point(0, 1, 0), Normalize(Vector(1, 1, 0))}
	e := Environment{Vector(0, -0.1, 0), Vector(-0.01, 0, 0)}
	for p.position.y > 0 {
		p = tick(e, p)
		fmt.Printf("Position: %v\n", p.position)
	}
}
