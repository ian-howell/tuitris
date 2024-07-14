package tetronimo

type T struct {
	grid [][]rune
}

func NewT() *T {
	return &T{
		grid: [][]rune{
			[]rune("   "),
			[]rune("TTT"),
			[]rune(" T "),
		},
	}
}

func (t T) Height() int {
	return 3
}

func (t T) Width() int {
	return 3
}

func (t T) Get(row, col int) rune {
	return t.grid[row][col]
}
