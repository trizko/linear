package linear

import (
	"reflect"
	"testing"
)

func TestMatrixSub(t *testing.T) {
	//should properly subtract two 2x1 matrices
	testmat1a := [][]float64{
		[]float64{1},
		[]float64{2},
	}
	testmat1b := [][]float64{
		[]float64{1},
		[]float64{2},
	}
	result1a := [][]float64{
		[]float64{0},
		[]float64{0},
	}
	cases1 := []struct {
		mat1 [][]float64
		mat2 [][]float64
		want [][]float64
	}{
		{testmat1a, testmat1b, result1a},
	}
	for _, c := range cases1 {
		got, _ := MatrixSub(c.mat1, c.mat2)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got: %v | wanted: %v", got, c.want)
		}
	}

	//should properly subtract three 2x1 matrices
	testmat2a := [][]float64{
		[]float64{1},
		[]float64{2},
	}
	testmat2b := [][]float64{
		[]float64{1},
		[]float64{2},
	}
	testmat2c := [][]float64{
		[]float64{1},
		[]float64{2},
	}
	result2a := [][]float64{
		[]float64{-1},
		[]float64{-2},
	}
	cases2 := []struct {
		mat1 [][]float64
		mat2 [][]float64
		mat3 [][]float64
		want [][]float64
	}{
		{testmat2a, testmat2b, testmat2c, result2a},
	}
	for _, c := range cases2 {
		got, _ := MatrixSub(c.mat1, c.mat2, c.mat3)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got: %v | wanted: %v", got, c.want)
		}
	}

	//should return an error for mismatched rows
	testmat3a := [][]float64{
		[]float64{1},
		[]float64{2},
	}
	testmat3b := [][]float64{
		[]float64{1},
		[]float64{2},
		[]float64{2},
	}
	cases3 := []struct {
		mat1 [][]float64
		mat2 [][]float64
	}{
		{testmat3a, testmat3b},
	}
	for _, c := range cases3 {
		_, err := MatrixSub(c.mat1, c.mat2)
		if err == nil {
			t.Errorf("error should not be nil because dimensions of %v do not match %v", c.mat1, c.mat2)
		}
	}

	//should return an error for mismatched columns
	testmat4a := [][]float64{
		[]float64{1, 2},
		[]float64{2, 3},
	}
	testmat4b := [][]float64{
		[]float64{1},
		[]float64{2},
	}
	cases4 := []struct {
		mat1 [][]float64
		mat2 [][]float64
	}{
		{testmat4a, testmat4b},
	}
	for _, c := range cases4 {
		_, err := MatrixSub(c.mat1, c.mat2)
		if err == nil {
			t.Errorf("error should not be nil because dimensions of %v do not match %v", c.mat1, c.mat2)
		}
	}

	//should return an error for mismatched rows and columns
	testmat5a := [][]float64{
		[]float64{1, 2},
		[]float64{2, 3},
	}
	testmat5b := [][]float64{
		[]float64{1},
		[]float64{2},
		[]float64{2},
	}
	cases5 := []struct {
		mat1 [][]float64
		mat2 [][]float64
	}{
		{testmat5a, testmat5b},
	}
	for _, c := range cases5 {
		_, err := MatrixSub(c.mat1, c.mat2)
		if err == nil {
			t.Errorf("error should not be nil because dimensions of %v do not match %v", c.mat1, c.mat2)
		}
	}

	//should properly subtract three 4x3 matrices
	testmat6a := [][]float64{
		[]float64{1, 2, 3},
		[]float64{1, 2, 3},
		[]float64{1, 2, 3},
		[]float64{1, 2, 3},
	}
	testmat6b := [][]float64{
		[]float64{1, 2, 3},
		[]float64{1, 2, 3},
		[]float64{1, 2, 3},
		[]float64{1, 2, 3},
	}
	testmat6c := [][]float64{
		[]float64{1, 2, 3},
		[]float64{1, 2, 3},
		[]float64{1, 2, 3},
		[]float64{1, 2, 3},
	}
	result6a := [][]float64{
		[]float64{-1, -2, -3},
		[]float64{-1, -2, -3},
		[]float64{-1, -2, -3},
		[]float64{-1, -2, -3},
	}
	cases6 := []struct {
		mat1 [][]float64
		mat2 [][]float64
		mat3 [][]float64
		want [][]float64
	}{
		{testmat6a, testmat6b, testmat6c, result6a},
	}
	for _, c := range cases6 {
		got, _ := MatrixSub(c.mat1, c.mat2, c.mat3)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got: %v | wanted: %v", got, c.want)
		}
	}
}
