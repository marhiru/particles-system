package particles

import (
	"fmt"
	"math"
	"time"
)

type Particle struct {
	Lifetime int64
	Speed    float64

	x float64
	y float64
}

type ParticleParams struct {
	MaxLife       int64
	MaxSpeed      float64
	ParticleCount int
	X             int
	Y             int
	nextPosition  NextPositionFunc
	ascii         Ascii
	reset         Reset
}

type NextPositionFunc func(particle *Particle, deltaMs int64)
type Ascii func(row, col int, count [][]int) rune
type Reset func(particle *Particle, params *ParticleParams) rune

type ParticleSystem struct {
	ParticleParams
	particles []*Particle

	lastTime int64
}

func NewParticleSystem(params ParticleParams, updateFunc NextPositionFunc) ParticleSystem {
	return ParticleSystem{
		ParticleParams: params,
		lastTime:       time.Now().UnixMilli(),
	}
}

func (ps *ParticleSystem) Start() {
	for _, p := range ps.particles {
		ps.reset(p, &ps.ParticleParams)
	}
}

func (ps *ParticleSystem) Update() {
	now := time.Now().UnixMilli()
	delta := now - ps.lastTime
	ps.lastTime = now

	for _, p := range ps.particles {
		ps.nextPosition(p, delta)

		if p.y >= float64(ps.Y) || p.x >= float64(ps.X) {
			ps.reset(p, &ps.ParticleParams)
		}
	}
}

func (ps *ParticleSystem) Print() [][]rune {
	counts := make([][]int, 0)

	for row := 0; row < ps.Y; row++ {
		count := make([]int, 0)
		for col := 0; col < ps.X; col++ {
			count = append(count, 0)
		}
		counts = append(counts, count)
	}

	activeParticles := 0
	for _, p := range ps.particles {
		row := int(math.Floor(p.y))
		col := int(math.Floor(p.x))

		if p.Lifetime > 0 && row >= 0 && row < len(counts) && col >= 0 && col < len(counts[row]) {
			activeParticles++
			fmt.Printf("DEBUG: Partícula ativa - x: %.2f, y: %.2f, row: %d, col: %d\n", p.x, p.y, row, col)
			counts[row][col]++
		}
	}

	fmt.Printf("DEBUG: Active particles - %d\n", activeParticles)

	out := make([][]rune, 0)
	for r, row := range counts {
		outRow := make([]rune, 0)
		for c := range row {
			rune := ps.ascii(c, r, counts)
			outRow = append(outRow, rune)

			if rune != ' ' {
				fmt.Printf("DEBUG: Rune - %c\n", rune)
			}
		}
		out = append(out, outRow)
	}
	return out
}
