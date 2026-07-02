package blocks

import "image/color"

// Shapes - Tetrisning barcha 7 ta klassik figuralari matrisasi
var Shapes = [][][]int{
	// 1. I Shakli (Havorang / Cyan)
	{
		{1, 1, 1, 1},
	},
	// 2. O Shakli (Sariq / Yellow)
	{
		{1, 1},
		{1, 1},
	},
	// 3. T Shakli (Siyohrang / Purple)
	{
		{0, 1, 0},
		{1, 1, 1},
	},
	// 4. S Shakli (Yashil / Green)
	{
		{0, 1, 1},
		{1, 1, 0},
	},
	// 5. Z Shakli (Qizil / Red)
	{
		{1, 1, 0},
		{0, 1, 1},
	},
	// 6. J Shakli (To'q ko'k / Blue)
	{
		{1, 0, 0},
		{1, 1, 1},
	},
	// 7. L Shakli (Olovrang / Orange)
	{
		{0, 0, 1},
		{1, 1, 1},
	},
}

// Colors - Har bir figuraga mos keluvchi rasmiy ranglar palette-si
// Ketma-ketlik yuqoridagi Shapes indeksi bilan bir xil bo'lishi shart
var Colors = []color.RGBA{
	{0, 255, 255, 255},   // I (Cyan)
	{255, 255, 0, 255},   // O (Yellow)
	{128, 0, 128, 255},   // T (Purple)
	{0, 255, 0, 255},     // S (Green)
	{255, 0, 0, 255},     // Z (Red)
	{0, 0, 255, 255},     // J (Blue)
	{255, 165, 0, 255},   // L (Orange)
}