package HRD

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

const (
	VERTICAL    = 0x01
	HORIZONTAL  = 0x02
	HORIZONTAL2 = 0x03
	BLOCK       = 0x04
	BLANK       = 0x05
	SQUARE      = 0x07
	SQUARE2     = 0x08
	ROWSIZE     = 5
	COLUMNSIZE  = 4
)

type StateNode struct {
	Pre   *StateNode
	State int64
}

var stateMap = make(map[int64]bool, 0)

func ItoM(n int64) [][]int64 {
	res := make([][]int64, 0)
	values := make([]int64, 0)
	idx := 0
	mask := int64(0x07)
	for ; idx < 20; idx++ {
		values = append(values, n&mask)
		n >>= 3
	}
	idx = 0
	for i := 0; i < 5; i++ {
		tmp := make([]int64, 4)
		for j := 0; j < 4; j++ {
			tmp[j] = values[idx]
			idx++
		}
		res = append(res, tmp)
	}
	return res
}

func MtoI(m [][]int64) int64 {
	var res int64
	idx := 0
	for _, r := range m {
		for _, v := range r {
			v <<= idx * 3
			res += v
			idx++
		}
	}
	return res
}

func StoM(s string) [][]int64 {
	res := make([][]int64, 0)
	values := make([]int64, 20)
	for i, c := range s {
		values[i] = int64(c - 48)
	}
	idx := 0
	for i := 0; i < 5; i++ {
		tmp := make([]int64, 4)
		for j := 0; j < 4; j++ {
			tmp[j] = values[idx]
			idx++
		}
		res = append(res, tmp)
	}
	return res
}

func PrintMv1(m [][]int64) {
	ui := [ROWSIZE][COLUMNSIZE]int{}
	for i, r := range m {
		for j, v := range r {
			switch v {
			case SQUARE:
				ui[i][j] = SQUARE
				ui[i][j+1] = SQUARE2
				ui[i+1][j] = SQUARE
				ui[i+1][j+1] = SQUARE2
			case BLOCK:
				ui[i][j] = BLOCK
			case VERTICAL:
				ui[i][j] = VERTICAL
				ui[i+1][j] = VERTICAL
			case HORIZONTAL:
				ui[i][j] = HORIZONTAL
				ui[i][j+1] = HORIZONTAL2
			case BLANK:
				ui[i][j] = BLANK
			}
		}
	}
	green := color.New(color.BgGreen)
	red := color.New(color.BgRed)
	white := color.New(color.BgHiWhite)
	yellow := color.New(color.BgYellow)
	cyan := color.New(color.BgHiCyan)
	for _, r := range ui {
		for _, v := range r {
			switch v {
			case SQUARE:
				green.Printf("曹")
				green.Printf(" ")
			case SQUARE2:
				green.Printf("曹")
				fmt.Printf(" ")
			case VERTICAL:
				red.Printf("竖")
				fmt.Printf(" ")
			case HORIZONTAL:
				cyan.Printf("横")
				cyan.Printf(" ")
			case HORIZONTAL2:
				cyan.Printf("横")
				fmt.Printf(" ")
			case BLOCK:
				yellow.Printf("兵")
				fmt.Printf(" ")
			case BLANK:
				white.Printf("口")
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

}

func PrintMv2(m [][]int64) {
	ui := [ROWSIZE][COLUMNSIZE]string{}
	for i, r := range m {
		for j, v := range r {
			switch v {
			case SQUARE:
				ui[i][j] = "┏"
				ui[i][j+1] = "┓"
				ui[i+1][j] = "┗"
				ui[i+1][j+1] = "┛"
			case BLOCK:
				ui[i][j] = "■"
			case VERTICAL:
				ui[i][j] = "ㄇ"
				ui[i+1][j] = "ㄩ"
			case HORIZONTAL:
				ui[i][j] = "ㄈ"
				ui[i][j+1] = "コ"
			case BLANK:
				ui[i][j] = "□"
			}
		}
	}
	for _, r := range ui {
		log.Println(r)
	}
}

func BlockMove(m [][]int64, x, y int, pre *StateNode) []*StateNode {
	res := make([]*StateNode, 0)
	if x-1 >= 0 && m[x-1][y] == BLANK {
		if node := MapSwitch(m, []int{x, y}, []int{x - 1, y}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
		if x-2 >= 0 && m[x-2][y] == BLANK {
			if node := MapSwitch(m, []int{x, y}, []int{x - 2, y}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
		if y-1 >= 0 && m[x-1][y-1] == BLANK {
			if node := MapSwitch(m, []int{x, y}, []int{x - 1, y - 1}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
		if y+1 < COLUMNSIZE && m[x-1][y+1] == BLANK {
			if node := MapSwitch(m, []int{x, y}, []int{x - 1, y + 1}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
	}
	if x+1 < ROWSIZE && m[x+1][y] == BLANK {
		if node := MapSwitch(m, []int{x, y}, []int{x + 1, y}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
		if x+2 < ROWSIZE && m[x+2][y] == BLANK {
			if node := MapSwitch(m, []int{x, y}, []int{x + 2, y}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
		if y-1 >= 0 && m[x+1][y-1] == BLANK {
			if node := MapSwitch(m, []int{x, y}, []int{x + 1, y - 1}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
		if y+1 < COLUMNSIZE && m[x+1][y+1] == BLANK {
			if node := MapSwitch(m, []int{x, y}, []int{x + 1, y + 1}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
	}
	if y-1 >= 0 && m[x][y-1] == BLANK {
		if node := MapSwitch(m, []int{x, y}, []int{x, y - 1}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
		if y-2 >= 0 && m[x][y-2] == BLANK {
			if node := MapSwitch(m, []int{x, y}, []int{x, y - 2}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
		if x+1 < ROWSIZE && m[x+1][y-1] == BLANK {
			if node := MapSwitch(m, []int{x, y}, []int{x + 1, y - 1}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
		if x-1 >= 0 && m[x-1][y-1] == BLANK {
			if node := MapSwitch(m, []int{x, y}, []int{x - 1, y - 1}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
	}

	if y+1 < COLUMNSIZE && m[x][y+1] == BLANK {
		if node := MapSwitch(m, []int{x, y}, []int{x, y + 1}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
		if y+2 < COLUMNSIZE && m[x][y+2] == BLANK {
			if node := MapSwitch(m, []int{x, y}, []int{x, y + 2}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
		if x+1 < ROWSIZE && m[x+1][y+1] == BLANK {
			if node := MapSwitch(m, []int{x, y}, []int{x + 1, y + 1}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
		if x-1 >= 0 && m[x-1][y+1] == BLANK {
			if node := MapSwitch(m, []int{x, y}, []int{x - 1, y + 1}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
	}

	return res
}

func SquareMove(m [][]int64, x, y int, pre *StateNode) []*StateNode {
	res := make([]*StateNode, 0)
	if x-1 >= 0 && m[x-1][y] == BLANK && m[x-1][y+1] == BLANK {
		if node := MapSwitch(m, []int{x, y, x, y + 1, x + 1, y, x + 1, y + 1}, []int{x - 1, y, x - 1, y + 1, x, y, x, y + 1}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
	}
	if x+2 < ROWSIZE && m[x+2][y] == BLANK && m[x+2][y+1] == BLANK {
		if node := MapSwitch(m, []int{x + 2, y, x + 2, y + 1, x + 1, y, x + 1, y + 1}, []int{x + 1, y, x + 1, y + 1, x, y, x, y + 1}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
	}
	if y-1 >= 0 && m[x][y-1] == BLANK && m[x+1][y-1] == BLANK {
		if node := MapSwitch(m, []int{x, y, x + 1, y, x, y + 1, x + 1, y + 1}, []int{x, y - 1, x + 1, y - 1, x, y, x + 1, y}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
	}
	if y+2 < COLUMNSIZE && m[x][y+2] == BLANK && m[x+1][y+2] == BLANK {
		if node := MapSwitch(m, []int{x, y + 1, x + 1, y + 1, x, y, x + 1, y}, []int{x, y + 2, x + 1, y + 2, x, y + 1, x + 1, y + 1}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
	}
	return res
}

func VerticalMove(m [][]int64, x, y int, pre *StateNode) []*StateNode {
	res := make([]*StateNode, 0)
	if x+2 < ROWSIZE && m[x+2][y] == BLANK {
		if node := MapSwitch(m, []int{x + 1, y, x, y}, []int{x + 2, y, x + 1, y}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
		if x+3 < ROWSIZE && m[x+3][y] == BLANK {
			if node := MapSwitch(m, []int{x + 1, y, x, y}, []int{x + 3, y, x + 2, y}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
	}
	if x-1 >= 0 && m[x-1][y] == BLANK {
		if node := MapSwitch(m, []int{x, y, x + 1, y}, []int{x - 1, y, x, y}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
		if x-2 >= 0 && m[x-2][y] == BLANK {
			if node := MapSwitch(m, []int{x, y, x + 1, y}, []int{x - 2, y, x - 1, y}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}

		}
	}
	if y+1 < COLUMNSIZE && m[x][y+1] == BLANK && m[x+1][y+1] == BLANK {
		if node := MapSwitch(m, []int{x, y, x + 1, y}, []int{x, y + 1, x + 1, y + 1}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
	}
	if y-1 >= 0 && m[x][y-1] == BLANK && m[x+1][y-1] == BLANK {
		if node := MapSwitch(m, []int{x, y, x + 1, y}, []int{x, y - 1, x + 1, y - 1}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
	}
	return res
}

func HorizontalMove(m [][]int64, x, y int, pre *StateNode) []*StateNode {
	res := make([]*StateNode, 0)
	if x-1 >= 0 && m[x-1][y] == BLANK && m[x-1][y+1] == BLANK {
		if node := MapSwitch(m, []int{x, y, x, y + 1}, []int{x - 1, y, x - 1, y + 1}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
	}
	if x+1 < ROWSIZE && m[x+1][y] == BLANK && m[x+1][y+1] == BLANK {
		if node := MapSwitch(m, []int{x, y, x, y + 1}, []int{x + 1, y, x + 1, y + 1}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
	}
	if y-1 >= 0 && m[x][y-1] == BLANK {
		if node := MapSwitch(m, []int{x, y, x, y}, []int{x, y - 1, x, y + 1}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
		if y-2 >= 0 && m[x][y-2] == BLANK {
			if node := MapSwitch(m, []int{x, y, x, y + 1}, []int{x, y - 2, x, y - 1}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
	}
	if y+2 < COLUMNSIZE && m[x][y+2] == BLANK {
		if node := MapSwitch(m, []int{x, y + 1, x, y}, []int{x, y + 2, x, y + 1}); node != nil {
			node.Pre = pre
			res = append(res, node)
		}
		if y+3 < COLUMNSIZE && m[x][y+3] == BLANK {
			if node := MapSwitch(m, []int{x, y, x, y + 1}, []int{x, y + 2, x, y + 3}); node != nil {
				node.Pre = pre
				res = append(res, node)
			}
		}
	}
	return res
}

func MapSwitch(m [][]int64, s []int, d []int) (node *StateNode) {
	for i := 0; i < len(s); {
		m[s[i]][s[i+1]], m[d[i]][d[i+1]] = m[d[i]][d[i+1]], m[s[i]][s[i+1]]
		i += 2
	}
	if isExist := AddToMap(m); !isExist {
		node = &StateNode{
			Pre:   nil,
			State: MtoI(m),
		}
	}
	for i := len(s) - 1; i >= 0; {
		m[s[i-1]][s[i]], m[d[i-1]][d[i]] = m[d[i-1]][d[i]], m[s[i-1]][s[i]]
		i -= 2
	}
	return
}

func AddToMap(m [][]int64) bool {
	tmp := MtoI(m)
	if _, ok := stateMap[tmp]; !ok {
		stateMap[tmp] = true
		return false
	}
	return true
}

func PrintStep(n *StateNode) {
	cnt := -1
	log.Println("********solution********")
	for n != nil {
		PrintMv1(ItoM(n.State))
		n = n.Pre
		cnt++
		fmt.Printf("\r--------%d steps last--------", cnt)
	}
	log.Println("********solution********")
	log.Printf("total steps %d (*^▽^*)", cnt)
}
