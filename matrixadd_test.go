package linear

import (
	"reflect"
	"testing"
)

func TestMatrixAdd(t *testing.T) {
	testmat1 := [][]float64{
		[]float64{1},
		[]float64{2},
	}
	testmat2 := [][]float64{
		[]float64{1},
		[]float64{2},
	}
	testmat3 := [][]float64{
		[]float64{1},
		[]float64{2},
	}
	testmat4 := [][]float64{
		[]float64{1, 2},
		[]float64{2, 3},
	}
	testmat5 := [][]float64{
		[]float64{1},
		[]float64{2},
		[]float64{2},
	}
	result1 := [][]float64{
		[]float64{2},
		[]float64{4},
	}
	result2 := [][]float64{
		[]float64{3},
		[]float64{6},
	}

	cases1 := []struct {
		mat1 [][]float64
		mat2 [][]float64
		want [][]float64
	}{
		{testmat1, testmat2, result1},
	}

	for _, c := range cases1 {
		got, _ := MatrixAdd(c.mat1, c.mat2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got: %v | wanted: %v", got, c.want)
		}
	}

	cases2 := []struct {
		mat1 [][]float64
		mat2 [][]float64
		mat3 [][]float64
		want [][]float64
	}{
		{testmat1, testmat2, testmat3, result2},
	}

	for _, c := range cases2 {
		got, _ := MatrixAdd(c.mat1, c.mat2, c.mat3)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got: %v | wanted: %v", got, c.want)
		}
	}

	casesNotMatchingDimensions := []struct {
		mat1 [][]float64
		mat2 [][]float64
	}{
		{testmat1, testmat4},
		{testmat1, testmat5},
	}

	for _, c := range casesNotMatchingDimensions {
		_, err := MatrixAdd(c.mat1, c.mat2)
		if err == nil {
			t.Errorf("error should not be nil because dimensions of %v do not match %v", c.mat1, c.mat2)
		}
	}
}
