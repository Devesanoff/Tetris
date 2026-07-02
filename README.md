# Go Tetris (Nostalgia)

A classic, nostalgic Tetris game built from scratch using **Go (Golang)** and the **Ebitengine (Ebiten v2)** 2D game framework. This project is developed using Clean Architecture (Modular Layout) principles to keep the game mechanics and visual assets cleanly separated.

> **"Just play and enjoy, my friend!"** ;)


---

##  Features
* **Pure Nostalgia:** Authentic arcade feel, mimicking the timeless handheld brick game mechanics.
* **7 Classic Tetrominoes:** Includes all traditional shapes (I, O, T, S, Z, J, L) with their official, standardized retro colors.
* **Modular Architecture:** Clean separation of concerns between block entities (`blocks/`) and core engine logic (`game/`).
* **Optimized Rendering:** Smooth vector-based rendering utilizing Ebitengine's native layout matrix.

---

##  Project Structure

```text
tetris/
├── blocks/
│   ├── block.go   # Block entity structure and dynamic positioning constructor
│   └── shapes.go  # Matrix definitions for the 7 Tetrominoes & color palettes
├── game/
│   └── game.go    # Game loop engine: handle inputs, gravity ticks, collision detection, and drawing
├── go.mod         # Go module definition
├── go.sum         # Dependency checksums
└── main.go        # Main entry point initializing the window and Ebitengine runner
```

---

##  Game Controls

| Key | Action |
| :--- | :--- |
| `⬅️ Left Arrow` | Move block Left |
| `➡️ Right Arrow` | Move block Right |
| `⬇️ Down Arrow` | Soft Drop (Speed up descent) |
| `⬆️ Up Arrow` / `Space` | Rotate Block 90° |

---

## Getting Started

### 1. Prerequisites (Ubuntu/Linux Setup)
Since Ebitengine relies on native graphics libraries via Cgo, ensure you have the necessary X11 and OpenGL development headers installed on your system:

```bash
sudo apt update && sudo apt install -y libgl1-mesa-dev libglu1-mesa-dev libx11-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxrender-dev libasound2-dev pkg-config libxxf86vm-dev
```

### 2. Clone and Run
Navigate into your project folder, clean up the modules, and run the main package:

```bash
# Tidy up Go modules
go mod tidy

# Run the game
go run main.go
```

---

##  Tech Stack
* **Language:** Go (Golang)
* **Game Engine:** [Ebitengine v2](https://github.com/hajimehoshi/ebiten)
* **OS Environment:** Linux (Ubuntu)
