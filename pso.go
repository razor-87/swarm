package main

import (
	"math"
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

type solution struct {
	pos []float64
	min float64
}

// Particle swarm optimization
func pso(f func([]float64) float64, min, max float64, dimensions int) (float64, []float64) {
	numCPU := runtime.NumCPU()
	swarms := numCPU - numCPU/4
	moves := swarms / 3
	pr := parameters[dimensions]
	pr.iterations /= numCPU
	halfMin := min / 2

	ring := make(chan solution, swarms)
	for n := 0; n < swarms; n++ {
		ring <- solution{pos: make([]float64, dimensions), min: math.MaxFloat64}
	}

	var wg sync.WaitGroup
	wg.Add(swarms)
	for c := 0; c < swarms; c++ {
		c := c
		go func() {
			defer wg.Done()

			var cSwarm [swarmSize]particle
			cBest := make([]float64, dimensions)
			cMin := math.MaxFloat64
			r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(c))) //#nosec G404

			for i := range cSwarm {
				cSwarm[i].vel, cSwarm[i].pos, cSwarm[i].pBest = make([]float64, dimensions), make([]float64, dimensions), make([]float64, dimensions)
				for d := 0; d < dimensions; d++ {
					cSwarm[i].vel[d] = max*r.Float64() + min
					cSwarm[i].pos[d] = max*r.Float64() + halfMin
					cSwarm[i].pBest[d] = cSwarm[i].pos[d]
				}
				if cSwarm[i].locMin = f(cSwarm[i].pos); cSwarm[i].locMin < cMin {
					cMin = cSwarm[i].locMin
					copy(cBest, cSwarm[i].pos)
				}
			}

			var sol solution
			for m := 0; m < moves; m++ {
				sol = <-ring
				if cMin < sol.min {
					sol.min = cMin
					copy(sol.pos, cBest)
				} else {
					cMin = sol.min
					copy(cBest, sol.pos)
				}
				ring <- sol
			}

			var chi, phiP, phiG, fitness float64
			for it := 0; it < pr.iterations; it++ {
				for i := range cSwarm {
					chi, phiP, phiG = r.Float64()*0.1+0.9, pr.cP*r.Float64(), pr.cG*r.Float64()
					for d := 0; d < dimensions; d++ {
						// vi = χ · (ω · vi + φ1 · (pi − xi) + φ2 · (pg − xi))
						cSwarm[i].vel[d] = chi * (pr.omega*cSwarm[i].vel[d] +
							phiP*(cSwarm[i].pBest[d]-cSwarm[i].pos[d]) +
							phiG*(cBest[d]-cSwarm[i].pos[d]))
						cSwarm[i].pos[d] += cSwarm[i].vel[d]
						switch {
						case cSwarm[i].pos[d] > max:
							cSwarm[i].pos[d] = max
						case cSwarm[i].pos[d] < min:
							cSwarm[i].pos[d] = min
						}
					}

					if fitness = f(cSwarm[i].pos); fitness < cSwarm[i].locMin {
						cSwarm[i].locMin = fitness
						copy(cSwarm[i].pBest, cSwarm[i].pos)

						if fitness < cMin {
							cMin = fitness
							copy(cBest, cSwarm[i].pos)
						}
					}
				}
			}

			sol = <-ring
			if cMin < sol.min {
				sol.min = cMin
				copy(sol.pos, cBest)
			}
			ring <- sol
		}()
	}
	wg.Wait()
	close(ring)

	gBest := make([]float64, dimensions)
	gMin := math.MaxFloat64
	for sol := range ring {
		if sol.min < gMin {
			gMin = sol.min
			copy(gBest, sol.pos)
		}
	}

	return gMin, gBest
}
