package particles

import (
	"math"
	"math/rand"
)

type CoffeeParticle struct {
	Particle
}

type CoffeeSystem struct {
	system    *ParticleSystem
	particles []*CoffeeParticle
}

func NewCoffeeSystem(width, height int) *CoffeeSystem {
	params := ParticleParams{
		MaxLife:       7,
		MaxSpeed:      0.5,
		ParticleCount: 100,
		X:             width,
		Y:             height,
	}

	system := NewParticleSystem(params, nil)
	system.particles = make([]*Particle, params.ParticleCount)

	system.nextPosition = coffeeNextPosition
	system.ascii = coffeeAscii
	system.reset = coffeeReset

	particles := make([]*CoffeeParticle, params.ParticleCount)

	for i := 0; i < params.ParticleCount; i++ {
		particles[i] = &CoffeeParticle{
			Particle: Particle{
				Lifetime: 0,
				Speed:    0,
				x:        0,
				y:        0,
			},
		}

		system.particles[i] = &particles[i].Particle
	}

	return &CoffeeSystem{
		system:    &system,
		particles: particles,
	}
}

func coffeeAscii(row, col int, count [][]int) rune {
	if row < 0 || row >= len(count) || col < 0 || col >= len(count[row]) {
		return ' '
	}

	if count[row][col] == 0 {
		return ' '
	}

	if count[row][col] == 1 {
		return '-'
	}

	if count[row][col] == 2 {
		return '}'
	}

	return '{'
}

func coffeeReset(particle *Particle, params *ParticleParams) rune {
	particle.Lifetime = int64(math.Floor(float64(params.MaxLife) * rand.Float64()))
	particle.Speed = math.Floor(params.MaxSpeed * rand.Float64())

	// maxXOffset := float64(params.X) / 2
	particle.x = math.Max(0, math.Min(rand.NormFloat64()*float64(params.X), float64(params.X)-1))
	particle.y = 0
	return '='
}

func coffeeNextPosition(particle *Particle, deltaMs int64) {
	particle.Lifetime -= deltaMs
	if particle.Lifetime < 0 {
		return
	}

	particle.y += particle.Speed * (float64(deltaMs) / 1000.0)
}

func (cs *CoffeeSystem) Start() {
	cs.system.Start()
}

func (cs *CoffeeSystem) Update() {
	cs.system.Update()
}

func (cs *CoffeeSystem) Print() [][]rune {
	return cs.system.Print()
}
