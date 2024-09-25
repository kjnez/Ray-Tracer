package features

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
	start := Point(0, 1, 0)
	velocity := Multiply(Normalize(Vector(1, 1.8, 0)), 11.25)
	p := Projectile{start, velocity}
	gravity := Vector(0, -0.1, 0)
	wind := Vector(-0.01, 0, 0)
	e := Environment{gravity, wind}
	canvas := *NewCanvas(900, 550)
	color := NewColor(0.5, 0.8, 0.2)
	for p.position.y >= 0 {
		p = tick(e, p)
		if p.position.y >= 0 {
			WritePixel(canvas, int(p.position.x), canvas.height - int(p.position.y), color)
		}
	}
	filename := "projectile.ppm"
	err := CanvasToPPM(canvas, filename)
	if err != nil {
		fmt.Printf("CanvasToPPM returned an error: %v", err)
	}
}
