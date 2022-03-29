package algorithms

//import "fmt"

// Inplace
// Swap element one by one
func rotate(matrix [][]int) {
	size := len(matrix)
	i, j := 0, 0
	countTotal, countRow := 0, 0
	curr := matrix[i][j]
	end := size - 1
	for {
		iNew := j
		jNew := size - 1 - i
		tmp := matrix[iNew][jNew]
		//fmt.Printf("%d. %d (%d,%d) -> %d (%d,%d)\n", countTotal, curr, i, j, tmp, iNew, jNew)
		matrix[iNew][jNew] = curr
		curr = tmp
		countTotal++
		countRow++
		i = iNew
		j = jNew
		if countTotal == size*size {
			break
		}
		if countRow%4 == 0 {
			j++
			if j == end {
				end--
				i++
				//fmt.Println("Next square")
				j = i
			}
			curr = matrix[i][j]
		}
	}
}
