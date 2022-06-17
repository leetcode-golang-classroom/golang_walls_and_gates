# golang_walls_and_gates

You are given a m x n 2D grid initialized with these three possible values.

- `-1`-  A wall or an obstacle.
- `0` - A gate.
- `INF` - Infinity means an empty room.
- We use the value $`2^{31}$ - 1 = 2147483647` to represent INF as you may assume that the distance to a gate is less than `2147483647`.Fill each empty room with the distance to its nearest gate. If it is impossible to reach a `Gate`, that room should remain filled with `INF`

*Contact me on wechat to get **Amazon、Google** requent Interview questions . (wechat id : **jiuzhang0607**)*

## Examples

**Example1**

```
Input:
[[2147483647,-1,0,2147483647],[2147483647,2147483647,2147483647,-1],[2147483647,-1,2147483647,-1],[0,-1,2147483647,2147483647]]
Output:
[[3,-1,0,1],[2,2,1,-1],[1,-1,2,-1],[0,-1,3,4]]

Explanation:
the 2D grid is:
INF  -1  0  INF
INF INF INF  -1
INF  -1 INF  -1
  0  -1 INF INF
the answer is:
  3  -1   0   1
  2   2   1  -1
  1  -1   2  -1
  0  -1   3   4

```

**Example2**

```
Input:
[[0,-1],[2147483647,2147483647]]
Output:
[[0,-1],[1,2]]

```

## 解析

類似於 [994. Rotting Oranges](https://www.notion.so/994-Rotting-Oranges-3c6396de56e74bef937c7038de36ea5f) 

題目給定了一個 m by n 整數矩陣 grid,

每個 grid[r][c] 有三種值

-1: 代表是一個牆或是阻礙物

0: 代表是一個門

INF： 這邊使用 $2^{31}$ -1 代表一個空房間

每個 cell 可以透過水平或垂直方向移動

實作一個演算法來找出每個 空房間到最近門的距離

類似的思路，可以先找出所有的門的座標

然後從門開始做對水平方向還有垂直方向做 BFS 更新每個 cell 的最短距離

直到 grid[r][c]  ≤ prevDist +1 就停止

否則更新 grid[r][c] = prevDist +1

![](https://i.imgur.com/YpQ0Xt4.png)


## 程式碼
```go
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
```
## 困難點

1. 理解如何透過BFS 累計目前到最近門的距離
2. 知道終止條件

## Solve Point

- [x]  理解如何透過BFS 累計目前到最近門的距離
- [x]  知道終止條件