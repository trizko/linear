// Copyright 20015 Tony Rizko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package linear contains functions that perform
// linear algebra operations
package linear

// ScalarMult multiplies a scalar number and matrix and returns the product
func ScalarMult(scalar float64, matrix [][]float64) [][]float64 {
	var result [][]float64

	for i, row := range matrix {
		if i == 0 {
			result = matrix
		}
		for j, val := range row {
			result[i][j] = scalar * val
		}
	}

	return result
}
