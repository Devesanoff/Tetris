package blocks

import "image/color"

// Block - Tushayotgan bitta figurani ifodalovchi struktura
type Block struct {
	Shape [][]int
	X, Y  int
	Color color.RGBA
}

// NewBlock - Tasodifiy tanlangan shakl va rang bilan yangi blok obyektini yaratadi
func NewBlock(shape [][]int, color color.RGBA, cols int) *Block {
	return &Block{
		Shape: shape,
		X:     cols/2 - len(shape[0])/2, // Maydon ustunlari sonidan kelib chiqib o'rtaga joylash
		Y:     0,
		Color: color,
	}
}