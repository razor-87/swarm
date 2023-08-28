package main

import "math"

const tau = 2 * math.Pi

// https://www.sfu.ca/~ssurjano/ackley.html
func ackley(xs []float64) float64 {
	var ss, sc float64
	for i := range xs {
		ss += xs[i] * xs[i]
		sc += math.Cos(tau * xs[i])
	}
	id := 1 / float64(len(xs))
	ss *= id
	sc *= id

	return -20*math.Exp(-0.2*math.Sqrt(ss)) - math.Exp(sc) + 20 + math.E
}

// https://www.sfu.ca/~ssurjano/beale.html
func beale(xs []float64) float64 {
	mul1 := xs[0] * xs[1]
	mul2 := mul1 * xs[1]
	return math.Pow(1.5-xs[0]+mul1, 2) + math.Pow(2.25-xs[0]+mul2, 2) + math.Pow(2.625-xs[0]+mul2*xs[1], 2)
}

// https://www.sfu.ca/~ssurjano/booth.html
func booth(xs []float64) float64 {
	return math.Pow(xs[0]+2*xs[1]-7, 2) + math.Pow(2*xs[0]+xs[1]-5, 2)
}

// https://www.sfu.ca/~ssurjano/drop.html
func dropWave(xs []float64) float64 {
	sqSum := xs[0]*xs[0] + xs[1]*xs[1]
	return -(1 + math.Cos(12*math.Sqrt(sqSum))) / (0.5*sqSum + 2)
}

// https://www.sfu.ca/~ssurjano/egg.html
func eggholder(xs []float64) float64 {
	z := xs[1] + 47
	return -z*math.Sin(math.Sqrt(math.Abs(xs[0]/2+z))) -
		xs[0]*math.Sin(math.Sqrt(math.Abs(xs[0]-z)))
}

// https://www.sfu.ca/~ssurjano/griewank.html
func griewank(xs []float64) float64 {
	var ss, pc float64
	for i := range xs {
		ss += xs[i] * xs[i]
		pc *= math.Cos(xs[i] / math.Sqrt(float64(i+1)))
	}

	return ss/4000 - (1 + pc) + 1
}

// https://www.sfu.ca/~ssurjano/holder.html
func holderTable(xs []float64) float64 {
	return -math.Abs(math.Sin(xs[0]) * math.Cos(xs[1]) *
		math.Exp(math.Abs(1-math.Sqrt(xs[0]*xs[0]+xs[1]*xs[1])/math.Pi)))
}

// https://www.sfu.ca/~ssurjano/levy.html
func levy(xs []float64) (ret float64) {
	lastIdx := len(xs) - 1
	for i := 0; i < lastIdx; i++ {
		wi := 1 + (xs[i]-1)/4
		s := math.Sin(math.Pi*wi + 1)
		ret += (wi - 1) * (wi - 1) * (1 + 10*s*s)
	}
	s1 := math.Sin(math.Pi * (1 + (xs[0]-1)/4))
	wd := 1 + (xs[lastIdx]-1)/4
	sd := math.Sin(tau * wd)

	return ret + s1*s1 + (wd-1)*(wd-1)*(1+sd*sd)
}

// https://www.sfu.ca/~ssurjano/rastr.html
func rastrigin(xs []float64) (ret float64) {
	for i := range xs {
		ret += xs[i]*xs[i] - 10*math.Cos(tau*xs[i]) + 10
	}
	return ret
}

// https://www.sfu.ca/~ssurjano/schaffer2.html
func schaffer2(xs []float64) float64 {
	sq0, sq1 := xs[0]*xs[0], xs[1]*xs[1]
	den, s := 1+0.001*(sq0+sq1), math.Sin(sq0-sq1)
	return 0.5 + (s*s-0.5)/(den*den)
}

// https://www.sfu.ca/~ssurjano/schwef.html
func schwefel(xs []float64) (ret float64) {
	for i := range xs {
		ret += xs[i] * math.Sin(math.Sqrt(math.Abs(xs[i])))
	}
	return 418.9829*float64(len(xs)) - ret
}

// https://www.sfu.ca/~ssurjano/stybtang.html
func styblinskiTang(xs []float64) (ret float64) {
	var sq float64
	for i := range xs {
		sq = xs[i] * xs[i]
		ret += sq*sq - 16*sq + 5*xs[i]
	}

	return ret / 2
}
