package linear

import (
	"reflect"
	"testing"
)

func TestScalarDiv(t *testing.T) {
	scalar1 := 3.
	matrix1 := [][]float64{
		[]float64{6},
		[]float64{9},
	}
	result1 := [][]float64{
		[]float64{2},
		[]float64{3},
	}
	scalar2 := 1.
	matrix2 := [][]float64{
		[]float64{1, 4},
		[]float64{2, 5},
	}
	result2 := [][]float64{
		[]float64{1, 4},
		[]float64{2, 5},
	}
	scalar3 := 5.
	matrix3 := [][]float64{
		[]float64{0, 4},
		[]float64{2, 5},
	}
	result3 := [][]float64{
		[]float64{0, .8},
		[]float64{.4, 1},
	}
	scalar4 := .5
	matrix4 := [][]float64{
		[]float64{1, 4},
		[]float64{2, 5},
	}
	result4 := [][]float64{
		[]float64{2, 8},
		[]float64{4, 10},
	}
	scalar5 := 5.
	matrix5 := [][]float64{
		[]float64{1, 4},
	}
	result5 := [][]float64{
		[]float64{.2, .8},
	}
	scalar6 := .5
	matrix6 := [][]float64{
		[]float64{1},
	}
	result6 := [][]float64{
		[]float64{2},
	}
	scalar7 := 1.5
	matrix7 := [][]float64{
		[]float64{4.5, 3, 4.5, 6, 9},
		[]float64{9, 3, 6, 3, 12},
	}
	result7 := [][]float64{
		[]float64{3, 2, 3, 4, 6},
		[]float64{6, 2, 4, 2, 8},
	}
	scalar8 := -10.
	matrix8 := [][]float64{
		[]float64{10, 30, 5, 8, -20},
		[]float64{-9, 3, -1, -2, 10},
		[]float64{9, 4, -1, 21, 1},
	}
	result8 := [][]float64{
		[]float64{-1, -3, -.5, -.8, 2},
		[]float64{.9, -.3, .1, .2, -1},
		[]float64{-.9, -.4, .1, -2.1, -.1},
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
	}

	for _, c := range cases {
		got, _ := ScalarDiv(c.scalar, c.matrix)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got: %v | wanted: %v", got, c.want)
		}
	}

	scalar1a := 0.
	matrix1a := [][]float64{
		[]float64{1, 4},
		[]float64{2, 5},
	}
	matrix1b := [][]float64{
		[]float64{1, 4},
		[]float64{2, 5},
		[]float64{2, 5},
	}
	matrix1c := [][]float64{
		[]float64{1, 4},
	}
	casesDivideByZero := []struct {
		scalar float64
		matrix [][]float64
	}{
		{scalar1a, matrix1a},
		{scalar1a, matrix1b},
		{scalar1a, matrix1c},
	}

	for _, c := range casesDivideByZero {
		_, err := ScalarDiv(c.scalar, c.matrix)
		if err == nil {
			t.Error("dividing by zero should return an error")
		}
	}
}
