package particles

import (
	"math"
	"math/rand"
)

type Coffee struct {
	ParticleSystem
}

func ascii(row, col int, count [][]int) rune {
	return '}'
}
func reset(particle *Particle, params *ParticleParams) rune {
	particle.lifetime = int64(math.Floor(float64(params.MaxLife) * rand.Float64()))
	return '}'
}

func nextPosition(particle *Particle, deltaMs int64) {

}
func NewCoffee(width, height int) Coffee {

	return Coffee{
		ParticleSystem: NewParticleSystem(
			ParticleParams{
				MaxLife:       7,
				MaxSpeed:      0.5,
				ParticleCount: 100,
				reset:         reset,
				nextPosition:  nextPosition,
				ascii:         ascii,
			},
		),
	}
}
