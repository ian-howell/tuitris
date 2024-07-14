package tetronimo

type Tetronimo interface {
	Height() int
	Width() int
	Get(row, col int) rune
}
