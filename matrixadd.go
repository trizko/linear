// Copyright 20015 Tony Rizko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package linear contains tools to implement machine
// learning in a Go environment.
package linear

import "errors"

// MatrixAdd takes 2 or more matrices and performs addition
func MatrixAdd(matrices ...[][]float64) ([][]float64, error) {
	rows := len(matrices[0])
	cols := len(matrices[0][0])
	for _, m := range matrices {
		if rows != len(m) {
			return nil, errors.New("all matrices must have same amount of rows")
		}
		if cols != len(m[0]) {
			return nil, errors.New("all matrices must have same amount of columns")
		}
	}

	result := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]float64, cols)
	}

	for _, mat := range matrices {
		for i, r := range mat {
			for j := range r {
				result[i][j] = result[i][j] + mat[i][j]
			}
		}
	}

	return result, nil
}
