package game

import (
	"image/color"
	"math/rand"

	"tetris/blocks"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	Cols      = 10
	Rows      = 20
	BlockSize = 20
)

type Game struct {
	board       [Rows][Cols]color.RGBA
	current     *blocks.Block 
	tickCounter int
}

func New() *Game {
	g := &Game{}
	g.spawnPiece()
	return g
}

func (g *Game) spawnPiece() {
	idx := rand.Intn(len(blocks.Shapes))
	
	g.current = blocks.NewBlock(blocks.Shapes[idx], blocks.Colors[idx], Cols)
}

func (g *Game) checkCollision(b *blocks.Block, dX, dY int) bool {
	for r := range b.Shape {
		for c := range b.Shape[r] {
			if b.Shape[r][c] != 0 {
				newX := b.X + c + dX
				newY := b.Y + r + dY

				if newX < 0 || newX >= Cols || newY >= Rows {
					return true
				}
				if newY >= 0 && g.board[newY][newX].A != 0 {
					return true
				}
			}
		}
	}
	return false
}

func rotateMatrix(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	rotated := make([][]int, m)
	for i := range rotated {
		rotated[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rotated[j][n-1-i] = matrix[i][j]
		}
	}
	return rotated
}

func (g *Game) lockPiece() {
	for r := range g.current.Shape {
		for c := range g.current.Shape[r] {
			if g.current.Shape[r][c] != 0 {
				g.board[g.current.Y+r][g.current.X+c] = g.current.Color
			}
		}
	}
	g.clearLines()
	g.spawnPiece()
}

func (g *Game) clearLines() {
	for r := Rows - 1; r >= 0; r-- {
		full := true
		for c := 0; c < Cols; c++ {
			if g.board[r][c].A == 0 {
				full = false
				break
			}
		}
		if full {
			for y := r; y > 0; y-- {
				g.board[y] = g.board[y-1]
			}
			g.board[0] = [Cols]color.RGBA{}
			r++
		}
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if !g.checkCollision(g.current, -1, 0) {
			g.current.X--
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if !g.checkCollision(g.current, 1, 0) {
			g.current.X++
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if !g.checkCollision(g.current, 0, 1) {
			g.current.Y++
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		rotated := rotateMatrix(g.current.Shape)
		tempBlock := &blocks.Block{Shape: rotated, X: g.current.X, Y: g.current.Y}
		if !g.checkCollision(tempBlock, 0, 0) {
			g.current.Shape = rotated
		}
	}

	g.tickCounter++
	if g.tickCounter > 30 {
		g.tickCounter = 0
		if !g.checkCollision(g.current, 0, 1) {
			g.current.Y++
		} else {
			g.lockPiece()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 20, 255})

	
	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			if g.board[r][c].A != 0 {
				x := float32(c * BlockSize)
				y := float32(r * BlockSize)
				vector.DrawFilledRect(screen, x, y, float32(BlockSize-1), float32(BlockSize-1), g.board[r][c], false)
			}
		}
	}

	
	if g.current != nil {
		for r := range g.current.Shape {
			for c := range g.current.Shape[r] {
				if g.current.Shape[r][c] != 0 {
					x := float32((g.current.X + c) * BlockSize)
					y := float32((g.current.Y + r) * BlockSize)
					vector.DrawFilledRect(screen, x, y, float32(BlockSize-1), float32(BlockSize-1), g.current.Color, false)
				}
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 240, 400
}