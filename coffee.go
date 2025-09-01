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

	return &CoffeeSystem{
		system:    &ParticleSystem{ParticleParams: params},
		particles: make([]*CoffeeParticle, params.ParticleCount),
	}
}

func (c *Coffee) ascii(row, col int, count [][]int) rune {
	return '}'
}

func (c *Coffee) reset() rune {
	c.Particle.Lifetime = int64(math.Floor(float64(c.ParticleParams.MaxLife) * rand.Float64()))
	c.Particle.Speed = math.Floor(c.ParticleParams.MaxSpeed * rand.Float64())

	maxXOffset := math.Floor(float64(c.ParticleParams.X) / 2)
	c.Particle.x = math.Max(-maxXOffset, math.Min(rand.NormFloat64(), maxXOffset))
	c.Particle.y = 0
	return '}'
}

func (c *Coffee) nextPosition(deltaMs int64) {
	c.Particle.Lifetime -= deltaMs
	if c.Particle.Lifetime < 0 {
		return
	}

	c.Particle.y += c.Particle.Speed * (float64(deltaMs) / 1000.0)
}
