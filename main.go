package main

import (
	"math/rand"
	"strconv"
)

// Maze is a struct for holding not only the grid but also the size and
// the stack used by the algorithm
type Maze struct {
	grid  [][]int
	stack [][2]int
	size  int
}

func newMaze(size int) *Maze {
	grid := make([][]int, size)

	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
		for j := 0; j < size; j++ {
			grid[i][j] = 15
		}
	}

	stack := make([][2]int, 0)

	return &Maze{
		grid,
		stack,
		size,
	}
}

func (m *Maze) getNeighbours(x int, y int) (neighbours [][2]int) {
	neighbours = make([][2]int, 0)

	if x > 0 {
		neighbours = append(neighbours,
			[2]int{x - 1, y},
		)
	}

	if x < m.size-1 {
		neighbours = append(neighbours,
			[2]int{x + 1, y},
		)
	}

	if y > 0 {
		neighbours = append(neighbours,
			[2]int{x, y - 1},
		)
	}

	if y < m.size-1 {
		neighbours = append(neighbours,
			[2]int{x, y + 1},
		)
	}

	return
}

func (m *Maze) hasBeenVisited(coords [2]int) bool {
	return m.grid[coords[1]][coords[0]] >= 16
}

func leftPad(str string, pad string, length int) (padded string) {
	padded = str

	if len(str) < length {
		for i := 0; i < length-len(str); i++ {
			padded = pad + padded
		}
	}

	return
}

func main() {
	size := 10
	maze := newMaze(size)

	initialX := rand.Intn(size)
	initialY := rand.Intn(size)

	maze.grid[initialY][initialX] += 16 // Mark as visited putting the most significant bit (MSB) to 1

	maze.stack = append(maze.stack, [2]int{initialX, initialY})

	for length := len(maze.stack); length > 0; {
		// Pop the last element of the stack
		currCoords := maze.stack[length-1]
		maze.stack = maze.stack[:length-1]

		currX := currCoords[0]
		currY := currCoords[1]

		neighbours := maze.getNeighbours(currX, currY)

		for _, el := range neighbours {
			if !maze.hasBeenVisited(el) {
				maze.stack = append(maze.stack, currCoords)
				xDiff := currX - el[0]
				yDiff := currY - el[1]

				if xDiff < 0 { // We are in the right neighbour
					maze.grid[currY][currX] -= 2 // Remove right wall of the current cell
					maze.grid[el[1]][el[0]] -= 8 // Remove the left wall of the chosen neighbour
				}

				if xDiff > 0 { // We are in the left neighbour
					maze.grid[currY][currX] -= 8 // Remove left wall of the current cell
					maze.grid[el[1]][el[0]] -= 2 // Remove right wall of the chosen neighbour
				}

				if yDiff < 0 { // We are in the bottom neighbour
					maze.grid[currY][currX]--    // Remove bottom wall of the current cell
					maze.grid[el[1]][el[0]] -= 4 // Remove the top wall of the chosen neighbour
				}

				if yDiff > 0 { // We are in the top neighbour
					maze.grid[currY][currX] -= 4 // Remove top wall of the current cell
					maze.grid[el[1]][el[0]]--    // Remove the bottom wall of the chosen neighbour
				}

				maze.grid[el[1]][el[0]] += 16 // Mark the neighbour as visited

				maze.stack = append(maze.stack, [2]int{el[0], el[1]}) // Push the cosen cell into the stack

				break // Stop looking for neighbours
			}
		}

		length = len(maze.stack)
	}

	for i := 0; i < size*3; i++ {
		print("_")
	}

	println()

	for i := range maze.grid {
		for j := range maze.grid[i] {
			binary := leftPad(strconv.FormatInt(int64(maze.grid[i][j]-16), 2), "0", 4)

			if string(binary[0]) == "1" { // left wall
				if j > 0 { // There's a neighbour on the left
					leftBinary := leftPad(strconv.FormatInt(int64(maze.grid[i][j-1]-16), 2), "0", 4)
					if string(leftBinary[2]) != "1" { // Left neighbour doesn't already have a right wall
						print("|")
					} else {
						print("_")
					}
				} else {
					print("|")
				}
			} else {
				print("_")
			}

			if string(binary[3]) == "1" { // botom wall
				print("_")
			} else {
				print(" ")
			}

			if string(binary[2]) == "1" { // right wall
				print("|")
			} else {
				print("_")
			}
		}
		println("")
	}
}
