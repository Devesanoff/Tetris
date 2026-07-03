package blocks

import "image/color"


type Block struct {
	Shape [][]int
	X, Y  int
	Color color.RGBA
}

func NewBlock(shape [][]int, color color.RGBA, cols int) *Block {
	return &Block{
		Shape: shape,
		X:     cols/2 - len(shape[0])/2, 
		Y:     0,
		Color: color,
	}
}