// 36. 有效的数独
// 判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

// 上图是一个部分填充的有效的数独。

// 数独部分空格内已填入了数字，空白格用 '.' 表示。

// 示例 1:

// 输入:
// [
//   ["5","3",".",".","7",".",".",".","."],
//   ["6",".",".","1","9","5",".",".","."],
//   [".","9","8",".",".",".",".","6","."],
//   ["8",".",".",".","6",".",".",".","3"],
//   ["4",".",".","8",".","3",".",".","1"],
//   ["7",".",".",".","2",".",".",".","6"],
//   [".","6",".",".",".",".","2","8","."],
//   [".",".",".","4","1","9",".",".","5"],
//   [".",".",".",".","8",".",".","7","9"]
// ]
// 输出: true
// 示例 2:

// 输入:
// [
//   ["8","3",".",".","7",".",".",".","."],
//   ["6",".",".","1","9","5",".",".","."],
//   [".","9","8",".",".",".",".","6","."],
//   ["8",".",".",".","6",".",".",".","3"],
//   ["4",".",".","8",".","3",".",".","1"],
//   ["7",".",".",".","2",".",".",".","6"],
//   [".","6",".",".",".",".","2","8","."],
//   [".",".",".","4","1","9",".",".","5"],
//   [".",".",".",".","8",".",".","7","9"]
// ]
// 输出: false
// 解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
//      但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。
// 说明:

// 一个有效的数独（部分已被填充）不一定是可解的。
// 只需要根据以上规则，验证已经填入的数字是否有效即可。
// 给定数独序列只包含数字 1-9 和字符 '.' 。
// 给定数独永远是 9x9 形式的。
package main

import "fmt"

func main() {
	Tboard := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	isValid := isValidSudoku(Tboard)
	fmt.Println(isValid)

}

func isValidSudoku(board [][]byte) bool {
	// fmt.Printf("%c\n", board[0][0])
	// var gh string = "5"
	// tempM := make(map[int]bool)

	//row col
	for i := 0; i < 9; i++ {
		tempR := make(map[byte]int, 10)
		tempC := make(map[byte]int, 10)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				tempR[board[i][j]]++
			}
		}
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				tempC[board[j][i]]++
			}
		}
		for _, v := range tempR {
			if v > 1 {
				return false
			}
		}
		for _, v := range tempC {
			if v > 1 {
				return false
			}
		}

	}

	//九块
	var n int = 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tempM := make(map[byte]int, 10)
			for row := i * n; row < (i+1)*n; row++ {
				for col := j * n; col < (j+1)*n; col++ {
					if board[row][col] != '.' {
						tempM[board[row][col]]++
					}
				}
			}
			for _, v := range tempM {
				if v > 1 {
					return false
				}
			}
		}
	}

	return true
}
