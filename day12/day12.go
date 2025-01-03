package day12

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/loissascha/go-assert/assert"
)

type PlotLine struct {
	char   string
	fields []int
	y      int
}

type CombinedPlotLine struct {
	char string
	rows []PlotLine
}

var maxX = 0
var maxY = 0

func Day12() {
	m := readFile("day12.input")
	maxY = len(m)
	for _, v := range m {
		maxX = len(v)
		fmt.Println(v)
	}
	combinedPlotLines := []CombinedPlotLine{}
	for y, line := range m {
		plotLines := mapLine(line, y)
		combinedPlotLines = combinePlotLines(combinedPlotLines, plotLines)
	}
	sum := 0
	for _, v := range combinedPlotLines {
		fmt.Println(v)
		per := printCombinedPlotLine(v)
		perRaw := calculatePerimeterString(per)
		perimeter := countPer(perRaw)
		perimeter2 := calculateSurrounding(perRaw)
		perimeter3 := countPerimeter3(perRaw)
		fields := countFields(perRaw)
		fmt.Println("perimeter:", perimeter)
		fmt.Println("perimeter2:", perimeter2)
		fmt.Println("perimeter3:", perimeter3)
		fmt.Println("fields:", fields)
		sum += (perimeter2 * fields)
	}
	fmt.Println("Sum:", sum)
}

func countFields(input [][]string) int {
	fields := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			char := input[y][x]
			if char != "+" && char != " " && char != "|" && char != "-" {
				fields++
			}
		}
	}
	return fields
}

func countPer(input [][]string) int {
	per := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			char := input[y][x]
			if char == "+" {
				per++
			}
		}
	}
	return per
}

func countPerimeter3(input [][]string) int {
	per := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			char := input[y][x]
			if char == "-" {
				touchingFields := 0
				if y-1 >= 0 {
					topChar := input[y-1][x]
					if topChar != " " {
						touchingFields++
					}
				}
				if y+1 < len(input) {
					bottomChar := input[y+1][x]
					if bottomChar != " " {
						touchingFields++
					}
				}
				if touchingFields < 2 {
					per++
				}
			} else if char == "|" {
				touchingFields := 0
				if x-1 >= 0 {
					leftChar := input[y][x-1]
					if leftChar != " " {
						touchingFields++
					}
				}
				fmt.Println("x + 1 =", x+1, "len:", len(input[y]))
				if x+1 < len(input[y]) {
					rightChar := input[y][x+1]
					if rightChar != " " {
						touchingFields++
					}
				}
				if touchingFields < 2 {
					per++
				}
			}
		}
	}
	return per
}

func calculateSurrounding(input [][]string) int {
	minY := len(input)
	maxY := 0
	minX := len(input[0])
	maxX := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			char := input[y][x]
			if char == "+" {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}

			}
		}
	}
	xDiff := maxX - minX
	yDiff := maxY - minY
	perimeter := (xDiff * 2) + (yDiff * 2)
	return perimeter / 2
}

func calculatePerimeterString(input [][]string) [][]string {
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			char := input[y][x]
			if char == " " {
				continue
			}
			if char == "+" {
				continue
			}
			if char == "|" {
				continue
			}
			if char == "-" {
				continue
			}
			leftX := x - 2
			rightX := x + 2
			topY := y - 2
			bottomY := y + 2
			writeLeft := false
			writeRight := false
			writeTop := false
			writeBottom := false
			if leftX > 0 {
				left := input[y][leftX]
				if left == " " {
					writeLeft = true
				}
			} else {
				writeLeft = true
			}

			if rightX < len(input[y]) {
				right := input[y][rightX]
				if right == " " {
					writeRight = true
				}
			} else {
				writeRight = true
			}

			if topY > 0 {
				top := input[topY][x]
				if top == " " {
					writeTop = true
				}
			} else {
				writeTop = true
			}

			if bottomY < len(input) {
				bottom := input[bottomY][x]
				if bottom == " " {
					writeBottom = true
				}
			} else {
				writeBottom = true
			}

			if writeLeft {
				input[y-1][x-1] = "+"
				input[y+1][x-1] = "+"
				input[y][x-1] = "|"
			}

			if writeRight {
				input[y-1][x+1] = "+"
				input[y+1][x+1] = "+"
				input[y][x+1] = "|"
			}

			if writeTop {
				input[y-1][x+1] = "+"
				input[y-1][x-1] = "+"
				input[y-1][x] = "-"
			}

			if writeBottom {
				input[y+1][x+1] = "+"
				input[y+1][x-1] = "+"
				input[y+1][x] = "-"
			}
		}
	}

	for y := 0; y < len(input); y++ {
		line := ""
		for x := 0; x < len(input[y]); x++ {
			char := input[y][x]
			line += char
		}
		if strings.TrimSpace(line) != "" {
			fmt.Println(line)
		}
	}

	return input
}

func printCombinedPlotLine(cpl CombinedPlotLine) [][]string {
	res := [][]string{}
	for yy := -1; yy < maxY; yy++ {
		resline := []string{}
		emptyline := []string{}
		for xx := -1; xx < maxX; xx++ {
			foundX := false
			for _, pl := range cpl.rows {
				y := pl.y
				if y != yy {
					continue
				}
				for _, x := range pl.fields {
					if x != xx {
						continue
					}
					resline = append(resline, pl.char)
					resline = append(resline, " ")
					foundX = true
				}
			}
			if !foundX {
				resline = append(resline, " ")
				resline = append(resline, " ")
			}
			emptyline = append(emptyline, " ")
			emptyline = append(emptyline, " ")
		}
		res = append(res, resline)
		res = append(res, emptyline)
	}
	return res
}

func hasPos(cpl CombinedPlotLine, x int, y int) bool {
	for _, pl := range cpl.rows {
		yy := pl.y
		if yy != y {
			continue
		}
		for _, xx := range pl.fields {
			if xx == x {
				return true
			}
		}
	}

	return false
}

func combinePlotLines(combinedPlotLines []CombinedPlotLine, plotLines []PlotLine) []CombinedPlotLine {
	workedPlotLinesIndexes := []int{}
	for i, cpl := range combinedPlotLines {
		for ii, pl := range plotLines {
			if pl.char != cpl.char {
				continue
			}
			canCombine := false
			for _, plfield := range pl.fields {
				for _, row := range cpl.rows {
					for _, field := range row.fields {
						if field == plfield && (pl.y == row.y-1 || pl.y == row.y+1) {
							// should be connected because fields match
							canCombine = true
						}
					}
				}
			}
			if canCombine {
				combinedPlotLines[i].rows = append(combinedPlotLines[i].rows, pl)
				workedPlotLinesIndexes = append(workedPlotLinesIndexes, ii)
			}
		}
	}

	for i, pl := range plotLines {
		found := false
		for _, wp := range workedPlotLinesIndexes {
			if wp == i {
				found = true
			}
		}
		if found {
			continue
		}
		combinedPlotLines = append(combinedPlotLines, CombinedPlotLine{char: pl.char, rows: []PlotLine{pl}})
	}

	return combinedPlotLines
}

func mapLine(line []string, y int) []PlotLine {
	plotLines := []PlotLine{}
	prevChar := ""
	currentFields := []int{}
	for x, char := range line {
		if char == prevChar {
			currentFields = append(currentFields, x)
			continue
		}

		if len(currentFields) > 0 {
			plotLines = append(plotLines, PlotLine{char: prevChar, fields: currentFields, y: y})
		}
		prevChar = char
		currentFields = []int{x}
	}
	if len(currentFields) > 0 {
		plotLines = append(plotLines, PlotLine{char: prevChar, fields: currentFields, y: y})
	}
	// fmt.Println(plotLines)
	return plotLines
}

func readFile(filepath string) [][]string {
	file, err := os.Open(filepath)
	assert.Nil(err, "Can't open file")
	defer file.Close()

	res := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ls := []string{}
		for i := 0; i < len(line); i++ {
			char := line[i : i+1]
			ls = append(ls, char)
		}
		res = append(res, ls)
	}
	return res
}
