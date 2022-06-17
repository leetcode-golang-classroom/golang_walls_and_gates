package sol

type Pair struct {
	row, col int
}

func WallsAndGates(rooms [][]int) {
	ROW := len(rooms)
	COL := len(rooms[0])
	queue := []Pair{}
	directions := []Pair{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for row := 0; row < ROW; row++ {
		for col := 0; col < COL; col++ {
			if rooms[row][col] == 0 {
				queue = append(queue, Pair{row: row, col: col})
			}
		}
	}
	for len(queue) != 0 {
		qLen := len(queue)
		for idx := 0; idx < qLen; idx++ {
			top := queue[0]
			queue = queue[1:]
			for _, direction := range directions {
				shifted_row := top.row + direction.row
				shifted_col := top.col + direction.col
				if shifted_row < 0 || shifted_row >= ROW || shifted_col < 0 || shifted_col >= COL ||
					rooms[shifted_row][shifted_col] <= rooms[top.row][top.col]+1 {
					continue
				}
				rooms[shifted_row][shifted_col] = rooms[top.row][top.col] + 1
				queue = append(queue, Pair{row: shifted_row, col: shifted_col})
			}
		}
	}
}
