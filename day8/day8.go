package day8

import (
	"bufio"
	"fmt"
	"os"

	"github.com/loissascha/go-assert/assert"
)

type Point struct {
	x, y int
}

var antinodes = map[Point]struct{}{}

func Day8() {
	grid := readFile("day8.input")
	antennas := parseAntennnas(grid)
	fmt.Println("Antennas:", antennas)

	findAntinodes(grid, antennas)

	for x, g := range grid {
		for y := 0; y < len(g); y++ {
			char := g[y : y+1]
			hasAntinode := false
			for a := range antinodes {
				if a.x == x && a.y == y {
					hasAntinode = true
					break
				}
			}

			if hasAntinode {
				fmt.Print("#")
				continue
			}
			fmt.Print(char)
		}
		fmt.Print("\n")
	}

	fmt.Println("points:", len(antinodes))
}

func findAntinodes(grid []string, antennas map[string][]Point) {
	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := 0; j < len(points); j++ {
				if i == j {
					continue
				}
				p1, p2 := points[i], points[j]
				// fmt.Println("trying to find antinode point for", ant, "with points", p1, "and", p2)

				dx, dy := p1.x-p2.x, p1.y-p2.y
				if dx < 0 {
					dx *= -1
				}
				if dy < 0 {
					dy *= -1
				}
				// fmt.Println("distance x:", dx, "distance y:", dy)

				createAntinotesDir1(p1, p2, dx, dy, 0, 0, grid)
				createAntinotesDir2(p1, p2, dx, dy, 0, 0, grid)
			}
		}
	}
}

func createAntinotesDir2(p1 Point, p2 Point, dx int, dy int, offsetx int, offsety int, grid []string) {
	ax := 0
	if p2.x > p1.x {
		ax = p2.x + dx
	} else if p2.x < p1.x {
		ax = p2.x - dx
	}
	ax += offsetx

	ay := 0
	if p2.y > p1.y {
		ay = p2.y + dy
	} else if p2.y < p1.y {
		ay = p2.y - dy
	}
	ay += offsety

	a := Point{x: ax, y: ay}

	if inBounds(a, grid) {
		antinodes[a] = struct{}{}
		createAntinotesDir2(p1, p2, dx, dy, offsetx+dx, offsety+dy, grid)
	}
}

func createAntinotesDir1(p1 Point, p2 Point, dx int, dy int, offsetx int, offsety int, grid []string) {
	ax := 0
	if p1.x > p2.x {
		ax = p1.x + dx
	} else if p1.x < p2.x {
		ax = p1.x - dx
	}
	ax += offsetx

	ay := 0
	if p1.y > p2.y {
		ay = p1.y + dy
	} else if p1.y < p2.y {
		ay = p1.y - dy
	}
	ay += offsety

	a := Point{x: ax, y: ay}
	if inBounds(a, grid) {
		antinodes[a] = struct{}{}
		createAntinotesDir1(p1, p2, dx, dy, offsetx+dx, offsety+dy, grid)
	}
}

func drawGridWithAntinote(grid []string, ax int, ay int) {
	for y, v := range grid {
		for x := 0; x < len(v); x++ {
			char := v[x : x+1]
			if y == ay && x == ax {
				fmt.Print("#")
			} else {
				fmt.Print(char)
			}
		}
		fmt.Print("\n")
	}
}

func inBounds(p Point, grid []string) bool {
	return p.y >= 0 && p.y < len(grid) && p.x >= 0 && p.x < len(grid[p.y])
}

func parseAntennnas(grid []string) map[string][]Point {
	antennas := map[string][]Point{}
	for y, line := range grid {
		for x := 0; x < len(line); x++ {
			char := line[x : x+1]
			if char != "." && char != "#" {
				fmt.Println("found antenna for char:", char, "pos:", x, y)
				antennas[char] = append(antennas[char], Point{x, y})
			}
		}
	}
	return antennas
}

func readFile(filepath string) []string {
	file, err := os.Open(filepath)
	assert.Nil(err, "Can't read file")
	defer file.Close()

	// Read the input map from stdin
	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		grid = append(grid, line)
	}
	return grid
}
