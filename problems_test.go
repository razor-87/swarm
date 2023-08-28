//go:build problems

package main

import (
	"math"
	"testing"
)

func Test_problems(t *testing.T) {
	epsilon := 1e-4
	type args struct {
		f  func([]float64) float64
		xs []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Ackley",
			args: args{
				f:  ackley,
				xs: make([]float64, 5),
			},
		},
		{
			name: "Beale",
			args: args{
				f:  beale,
				xs: []float64{3.0, 0.5},
			},
		},
		{
			name: "Booth",
			args: args{
				f:  booth,
				xs: []float64{1.0, 3.0},
			},
		},
		{
			name: "Drop-Wave",
			args: args{
				f:  dropWave,
				xs: make([]float64, 2),
			},
			want: -1,
		},
		{
			name: "Eggholder",
			args: args{
				f:  eggholder,
				xs: []float64{512, 404.2319},
			},
			want: -959.64066,
		},
		{
			name: "Griewank",
			args: args{
				f:  griewank,
				xs: make([]float64, 5),
			},
		},
		{
			name: "Holder Table",
			args: args{
				f:  holderTable,
				xs: []float64{8.05502, 9.66459},
			},
			want: -19.2085,
		},
		{
			name: "Levy",
			args: args{
				f:  levy,
				xs: []float64{1, 1, 1, 1, 1},
			},
		},
		{
			name: "Rastrigin",
			args: args{
				f:  rastrigin,
				xs: make([]float64, 5),
			},
		},
		{
			name: "Schaffer N.2",
			args: args{
				f:  schaffer2,
				xs: make([]float64, 2),
			},
		},
		{
			name: "Schwefel",
			args: args{
				f:  schwefel,
				xs: []float64{420.9687, 420.9687, 420.9687, 420.9687, 420.9687},
			},
		},
		{
			name: "Styblinski-Tang",
			args: args{
				f:  styblinskiTang,
				xs: []float64{-2.903534, -2.903534, -2.903534, -2.903534, -2.903534},
			},
			want: -195.8308285, // -39.1661657 * dimensions
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.args.f(tt.args.xs); got != tt.want {
				if !(math.Abs(tt.want-got) < epsilon) {
					t.Errorf("f() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
