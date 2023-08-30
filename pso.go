package main

import (
	"math"
	"math/bits"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const swarmSize = 223

// M.E.H. Pedersen - Good Parameters for Particle Swarm Optimization (2010)
var parameters = [...]parameter{
	2: {4000, -0.2797, 1.5539, 3.0539},
	5: {10000, -0.3699, -0.1207, 3.3657},
}

type parameter struct {
	iterations    int
	omega, cP, cG float64
}

type particle struct {
	vel, pos, pBest []float64
	locMin          float64
}

// Particle swarm optimization
func pso(f func([]float64) float64, min, max float64, dimensions int) (float64, []float64) {
	pr := parameters[dimensions]
	numCPU := runtime.NumCPU()
	if numCPU&(numCPU-1) == 0 && numCPU < 256 {
		p := bits.TrailingZeros8(uint8(numCPU))
		pr.iterations >>= p << 1
		pr.iterations += p
	} else {
		pr.iterations /= numCPU * numCPU
		pr.iterations += int(math.Sqrt(float64(numCPU)))
	}

	signal := make(chan struct{})
	chBest := make(chan []float64)
	chMinimum := make(chan float64)
	go func() {
		var best []float64
		minimum := math.MaxFloat64
		for {
			select {
			case minimum = <-chMinimum:
			case chMinimum <- minimum:
			case best = <-chBest:
			case chBest <- best:
			case <-signal:
				return
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(numCPU)
	for c := 0; c < numCPU; c++ {
		c := c
		go func() {
			defer wg.Done()
			var (
				fitness float64
				swarm   [swarmSize]particle
				r       *rand.Rand
			)
			swarmBest := make([]float64, dimensions)
			swarmMinimum := math.MaxFloat64

			r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(c))) //#nosec G404
			for i := range swarm {
				swarm[i].vel, swarm[i].pos, swarm[i].pBest = make([]float64, dimensions), make([]float64, dimensions), make([]float64, dimensions)
				for d := 0; d < dimensions; d++ {
					swarm[i].vel[d] = max*r.Float64() + min
					swarm[i].pos[d] = max*r.Float64() + min/2
					swarm[i].pBest[d] = swarm[i].pos[d]
				}
				fitness = f(swarm[i].pos)
				swarm[i].locMin = fitness
				if fitness < swarmMinimum {
					copy(swarmBest, swarm[i].pos)
					swarmMinimum = fitness
				}
			}

			var chi, phiP, phiG float64
			syncs := r.Intn(numCPU) + 1
			if syncs == 1 {
				syncs += r.Intn(numCPU)
			}
			for s := 0; s < syncs; s++ {
				r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(syncs))) //#nosec G404
				for it := 0; it < pr.iterations; it++ {
					for i := range swarm {
						chi, phiP, phiG = r.Float64()*0.1+0.9, pr.cP*r.Float64(), pr.cG*r.Float64()
						for d := 0; d < dimensions; d++ {
							// vi = χ · (ω · vi + φ1 · (pi − xi) + φ2 · (pg − xi))
							swarm[i].vel[d] = chi * (pr.omega*swarm[i].vel[d] +
								phiP*(swarm[i].pBest[d]-swarm[i].pos[d]) +
								phiG*(swarmBest[d]-swarm[i].pos[d]))
							swarm[i].pos[d] += swarm[i].vel[d]
							switch {
							case swarm[i].pos[d] > max:
								swarm[i].pos[d] = max
							case swarm[i].pos[d] < min:
								swarm[i].pos[d] = min
							}
						}
						fitness = f(swarm[i].pos)
						if fitness < swarm[i].locMin {
							copy(swarm[i].pBest, swarm[i].pos)
							swarm[i].locMin = fitness
						}
						if fitness < swarmMinimum {
							copy(swarmBest, swarm[i].pos)
							swarmMinimum = fitness
						}
					}
				}
				if minimum := <-chMinimum; swarmMinimum < minimum {
					chMinimum <- swarmMinimum
					chBest <- swarmBest
				}
			}
		}()
	}
	wg.Wait()
	gMinimum, gBest := <-chMinimum, <-chBest
	signal <- struct{}{}

	return gMinimum, gBest
}
