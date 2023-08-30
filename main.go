package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	minimum, position := pso(ackley, -32.768, 32.768, 5)
	elapsed := time.Since(start)
	fmt.Printf("Global minimum: %.3f, position: %.3f\n", minimum, position)
	fmt.Printf("Measure time: %s\n", elapsed)
}
