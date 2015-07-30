package linear

import (
	"reflect"
	"testing"
)

func TestScalarMult(t *testing.T) {
	scalar1 := 3.
	matrix1 := [][]float64{
		[]float64{1},
		[]float64{2},
	}
	result1 := [][]float64{
		[]float64{3},
		[]float64{6},
	}
	scalar2 := 0.
	matrix2 := [][]float64{
		[]float64{1, 4},
		[]float64{2, 5},
	}
	result2 := [][]float64{
		[]float64{0, 0},
		[]float64{0, 0},
	}
	scalar3 := 5.
	matrix3 := [][]float64{
		[]float64{0, 4},
		[]float64{2, 5},
	}
	result3 := [][]float64{
		[]float64{0, 20},
		[]float64{10, 25},
	}
	scalar4 := .5
	matrix4 := [][]float64{
		[]float64{1, 4},
		[]float64{2, 5},
	}
	result4 := [][]float64{
		[]float64{.5, 2},
		[]float64{1, 2.5},
	}
	scalar5 := 5.
	matrix5 := [][]float64{
		[]float64{1, 4},
	}
	result5 := [][]float64{
		[]float64{5, 20},
	}
	scalar6 := .5
	matrix6 := [][]float64{
		[]float64{1},
	}
	result6 := [][]float64{
		[]float64{.5},
	}
	scalar7 := 1.5
	matrix7 := [][]float64{
		[]float64{1, 3, 4, 7, 9},
		[]float64{9, 3, 1, 2, 10},
		[]float64{9, 4, 1, 21, 1},
	}
	result7 := [][]float64{
		[]float64{1.5, 4.5, 6, 10.5, 13.5},
		[]float64{13.5, 4.5, 1.5, 3, 15},
		[]float64{13.5, 6, 1.5, 31.5, 1.5},
	}
	scalar8 := 10.
	matrix8 := [][]float64{
		[]float64{1, 3, 4, 7, 9},
		[]float64{9, 3, 1, 2, 10},
		[]float64{9, 4, 1, 21, 1},
	}
	result8 := [][]float64{
		[]float64{10, 30, 40, 70, 90},
		[]float64{90, 30, 10, 20, 100},
		[]float64{90, 40, 10, 210, 10},
	}
	scalar9 := -10.
	matrix9 := [][]float64{
		[]float64{1, 3, 4, 7},
		[]float64{9, 3, 1, 2},
		[]float64{9, 4, 1, 21},
	}
	result9 := [][]float64{
		[]float64{-10, -30, -40, -70},
		[]float64{-90, -30, -10, -20},
		[]float64{-90, -40, -10, -210},
	}
	scalar10 := -.5
	matrix10 := [][]float64{
		[]float64{1, 3},
		[]float64{9, 3},
		[]float64{10, -4},
	}
	result10 := [][]float64{
		[]float64{-.5, -1.5},
		[]float64{-4.5, -1.5},
		[]float64{-5, 2},
	}
	cases := []struct {
		scalar float64
		matrix [][]float64
		want   [][]float64
	}{
		{scalar1, matrix1, result1},
		{scalar2, matrix2, result2},
		{scalar3, matrix3, result3},
		{scalar4, matrix4, result4},
		{scalar5, matrix5, result5},
		{scalar6, matrix6, result6},
		{scalar7, matrix7, result7},
		{scalar8, matrix8, result8},
		{scalar9, matrix9, result9},
		{scalar10, matrix10, result10},
	}

	for _, c := range cases {
		got := ScalarMult(c.scalar, c.matrix)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got: %v | wanted: %v", got, c.want)
		}
	}
}
