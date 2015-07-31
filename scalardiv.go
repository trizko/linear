// Copyright 20015 Tony Rizko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package linear contains functions that perform
// linear algebra operations
package linear

import "errors"

// ScalarDiv divides a matrix by a scalar value and returns the quotient
func ScalarDiv(scalar float64, matrix [][]float64) ([][]float64, error) {
	var result [][]float64

	if scalar == 0 {
		return nil, errors.New("cannot divide by zero")
	}

	for i, row := range matrix {
		if i == 0 {
			result = matrix
		}
		for j, val := range row {
			result[i][j] = val / scalar
		}
	}

	return result, nil
}
