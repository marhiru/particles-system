package particles

import (
	"fmt"
)

func main() {
	coffee := NewCoffee(10)
	for _, p := range coffee.particles {
		fmt.Println(p.x, p.y)
	}
}
