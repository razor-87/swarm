//go:build !problems

package main

import (
	"math"
	"testing"
)

func Test_pso(t *testing.T) {
	epsilon := 1e-2
	type args struct {
		f          func([]float64) float64
		min        float64
		max        float64
		dimensions int
	}
	tests := []struct {
		name     string
		args     args
		want     float64
		want1    []float64
		want1Abs bool
	}{
		{
			name: "Ackley",
			args: args{
				f:          ackley,
				min:        -32.768,
				max:        32.768,
				dimensions: 5,
			},
			want1: make([]float64, 5),
		},
		{
			name: "Beale",
			args: args{
				f:          beale,
				min:        -4.5,
				max:        4.5,
				dimensions: 2,
			},
			want1: []float64{3.0, 0.5},
		},
		{
			name: "Booth",
			args: args{
				f:          booth,
				min:        -10,
				max:        10,
				dimensions: 2,
			},
			want1: []float64{1.0, 3.0},
		},
		{
			name: "Drop-Wave",
			args: args{
				f:          dropWave,
				min:        -5.12,
				max:        5.12,
				dimensions: 2,
			},
			want:  -1,
			want1: make([]float64, 2),
		},
		{
			name: "Eggholder",
			args: args{
				f:          eggholder,
				min:        -512,
				max:        512,
				dimensions: 2,
			},
			want:  -959.64066,
			want1: []float64{512, 404.2319},
		},
		{
			name: "Griewank",
			args: args{
				f:          griewank,
				min:        -600,
				max:        600,
				dimensions: 5,
			},
			want1: make([]float64, 5),
		},
		{
			name: "Holder Table",
			args: args{
				f:          holderTable,
				min:        -10,
				max:        10,
				dimensions: 2,
			},
			want:     -19.2085,
			want1:    []float64{8.05502, 9.66459},
			want1Abs: true,
		},
		{
			name: "Levy",
			args: args{
				f:          levy,
				min:        -10,
				max:        10,
				dimensions: 5,
			},
			want1: []float64{1, 1, 1, 1, 1},
		},
		{
			name: "Rastrigin",
			args: args{
				f:          rastrigin,
				min:        -5.12,
				max:        5.12,
				dimensions: 5,
			},
			want1: make([]float64, 5),
		},
		{
			name: "Schaffer N.2",
			args: args{
				f:          schaffer2,
				min:        -100,
				max:        100,
				dimensions: 2,
			},
			want1: make([]float64, 2),
		},
		{
			name: "Schwefel",
			args: args{
				f:          schwefel,
				min:        -500,
				max:        500,
				dimensions: 5,
			},
			want1: []float64{420.9687, 420.9687, 420.9687, 420.9687, 420.9687},
		},
		{
			name: "Styblinski-Tang",
			args: args{
				f:          styblinskiTang,
				min:        -5,
				max:        5,
				dimensions: 5,
			},
			want:  -195.8308285, // -39.1661657 * dimensions
			want1: []float64{-2.903534, -2.903534, -2.903534, -2.903534, -2.903534},
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, got1 := pso(tt.args.f, tt.args.min, tt.args.max, tt.args.dimensions)
			if !(math.Abs(tt.want-got) < epsilon) {
				t.Errorf("pso() got = %v, want %v", got, tt.want)
			}
			if tt.want1Abs {
				for i := range got1 {
					got1[i] = math.Abs(got1[i])
				}
			}
			for i := range got1 {
				if !(math.Abs(tt.want1[i]-got1[i]) < epsilon) {
					t.Errorf("pso() got1[%d] = %v, want %v", i, got1[i], tt.want1[i])
				}
			}
		})
	}
}
