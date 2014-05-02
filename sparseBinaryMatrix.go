package htm

import (
//"math"
)

//Entries are positions of non-zero values
type SparseEntry struct {
	Row int
	Col int
}

//Sparse binary matrix stores indexes of non-zero entries in matrix
//to conserve space
type SparseBinaryMatrix struct {
	Width             int
	Height            int
	TotalNonZeroCount int
	Entries           []SparseEntry
}

//Create new sparse binary matrix of specified size
func NewSparseBinaryMatrix(height, width int) *SparseBinaryMatrix {
	m := &SparseBinaryMatrix{}
	m.Height = height
	m.Width = width
	//Intialize with 70% sparsity
	//m.Entries = make([]SparseEntry, int(math.Ceil(width*height*0.3)))
	return m
}

//Create sparse binary matrix from specified dense matrix
func NewSparseBinaryMatrixFromDense(values [][]bool) *SparseBinaryMatrix {
	if len(values) < 1 {
		panic("No values specified.")
	}
	m := &SparseBinaryMatrix{}
	m.Height = len(values)
	m.Width = len(values[0])

	for r := 0; r < len(values); r++ {
		m.SetRowFromDense(r, values[r])
	}

	return m
}

// func NewRandSparseBinaryMatrix() *SparseBinaryMatrix {
// }

// func (sm *SparseBinaryMatrix) Resize(width int, height int) {
// }

//Get value at col,row position
func (sm *SparseBinaryMatrix) Get(row int, col int) bool {
	for _, val := range sm.Entries {
		if val.Row == row && val.Col == col {
			return true
		}
	}
	return false
}

func (sm *SparseBinaryMatrix) delete(row int, col int) {
	for idx, val := range sm.Entries {
		if val.Row == row && val.Col == col {
			sm.Entries = append(sm.Entries[:idx], sm.Entries[idx+1:]...)
			break
		}
	}

}

//Set value at row,col position
func (sm *SparseBinaryMatrix) Set(row int, col int, value bool) {
	if !value {
		sm.delete(row, col)
		return
	}

	if sm.Get(row, col) {
		return
	}

	newEntry := SparseEntry{}
	newEntry.Col = col
	newEntry.Row = row
	sm.Entries = append(sm.Entries, newEntry)

}

//Replaces specified row with values, assumes values is ordered
//correctly
func (sm *SparseBinaryMatrix) ReplaceRow(row int, values []bool) {
	sm.validateRowCol(row, len(values))

	for i := 0; i < sm.Width; i++ {
		sm.Set(row, i, values[i])
	}
}

//Replaces row with true values at specified indices
func (sm *SparseBinaryMatrix) ReplaceRowByIndices(row int, indices []int) {
	sm.validateRow(row)

	for i := 0; i < sm.Width; i++ {
		val := false
		for x := 0; x < len(indices); x++ {
			if i == indices[x] {
				val = true
				break
			}
		}
		sm.Set(row, i, val)
	}
}

//Returns dense row
func (sm *SparseBinaryMatrix) GetDenseRow(row int) []bool {
	sm.validateRow(row)
	result := make([]bool, sm.Width)

	for i := 0; i < len(sm.Entries); i++ {
		if sm.Entries[i].Row == row {
			result[sm.Entries[i].Col] = true
		}
	}

	return result
}

//Returns a rows "on" indices
func (sm *SparseBinaryMatrix) GetRowIndices(row int) []int {
	result := []int{}
	for i := 0; i < len(sm.Entries); i++ {
		if sm.Entries[i].Row == row {
			result = append(result, sm.Entries[i].Col)
		}
	}
	return result
}

//Sets a sparse row from dense representation
func (sm *SparseBinaryMatrix) SetRowFromDense(row int, denseRow []bool) {
	sm.validateRowCol(row, len(denseRow))
	for i := 0; i < sm.Width; i++ {
		sm.Set(i, row, denseRow[i])
	}
}

//In a normal matrix this would be multiplication in binary terms
//we just and then sum the true entries
func (sm *SparseBinaryMatrix) RowAndSum(row []bool) []int {
	sm.validateCol(len(row))
	result := make([]int, sm.Height)

	for _, val := range sm.Entries {
		if row[val.Col] {
			result[val.Col]++
		}
	}

	return result
}

func (sm *SparseBinaryMatrix) validateCol(col int) {
	if col > sm.Width {
		panic("Specified row is wider than matrix.")
	}
}

func (sm *SparseBinaryMatrix) validateRow(row int) {
	if row > sm.Height {
		panic("Specified row is out of bounds.")
	}
}

func (sm *SparseBinaryMatrix) validateRowCol(row int, col int) {
	sm.validateCol(col)
	sm.validateRow(row)
}
