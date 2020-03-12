package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

// Coord one coordinate
type Coord struct {
	x, y int
}

// TT one tetromino struct
type TT struct {
	name   string
	ttname string
}

var lib = map[string][4]Coord{
	"I":  [4]Coord{Coord{0, 0}, Coord{0, 1}, Coord{0, 2}, Coord{0, 3}},
	"IH": [4]Coord{Coord{0, 0}, Coord{1, 0}, Coord{2, 0}, Coord{3, 0}},
	"O":  [4]Coord{Coord{0, 0}, Coord{1, 0}, Coord{0, 1}, Coord{1, 1}},
	"L":  [4]Coord{Coord{0, 0}, Coord{0, 1}, Coord{0, 2}, Coord{1, 2}},
	"LR": [4]Coord{Coord{0, 0}, Coord{1, 0}, Coord{2, 0}, Coord{0, 1}},
	"LD": [4]Coord{Coord{0, 0}, Coord{1, 0}, Coord{1, 1}, Coord{1, 2}},
	"LL": [4]Coord{Coord{2, 0}, Coord{0, 1}, Coord{1, 1}, Coord{2, 1}},
	"J":  [4]Coord{Coord{1, 0}, Coord{1, 1}, Coord{0, 2}, Coord{1, 2}},
	"JR": [4]Coord{Coord{0, 0}, Coord{0, 1}, Coord{1, 1}, Coord{2, 1}},
	"JD": [4]Coord{Coord{0, 0}, Coord{1, 0}, Coord{0, 1}, Coord{0, 2}},
	"JL": [4]Coord{Coord{0, 0}, Coord{1, 0}, Coord{2, 0}, Coord{2, 1}},
	"T":  [4]Coord{Coord{1, 0}, Coord{0, 1}, Coord{1, 1}, Coord{2, 1}},
	"TR": [4]Coord{Coord{0, 0}, Coord{0, 1}, Coord{1, 1}, Coord{0, 2}},
	"TD": [4]Coord{Coord{0, 0}, Coord{1, 0}, Coord{2, 0}, Coord{1, 1}},
	"TL": [4]Coord{Coord{1, 0}, Coord{0, 1}, Coord{1, 1}, Coord{1, 2}},
	"S":  [4]Coord{Coord{1, 0}, Coord{2, 0}, Coord{0, 1}, Coord{1, 1}},
	"SR": [4]Coord{Coord{0, 0}, Coord{0, 1}, Coord{1, 1}, Coord{1, 2}},
	"Z":  [4]Coord{Coord{0, 0}, Coord{1, 0}, Coord{1, 1}, Coord{2, 1}},
	"ZR": [4]Coord{Coord{1, 0}, Coord{0, 1}, Coord{1, 1}, Coord{0, 2}},
}

func main() {

	filename := "test1.txt"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	str := strings.Split(strings.Trim(string(file), "\n \t\r"), "\n")
	tetromino := [][]string{[]string{}}
	x := 0
	for i, v := range str {
		if (i+1)%5 == 0 {
			x++
			tetromino = append(tetromino, []string{})
		}
		tetromino[x] = append(tetromino[x], v)
	}

	l := int(math.Ceil(2 * math.Sqrt(float64(len(tetromino)))))

	tts := getTetrominoes(tetromino)

	fmt.Println(tts)

	square := make([]string, l)
	for i := range square {
		for j := 0; j < l; j++ {
			square[i] += "."
		}
	}

	fillIt(tts, square)
}

func fillIt(tts []TT, square []string) {

	printSquare(square)

	if len(tts) == 0 {

		return
	}

	tt := tts[0]
	tts = tts[1:]

	fmt.Println(tts)

	// Outer:
	for y, line := range square {
		for x := range line {

			backtrack := make([]string, len(square))

			for i, item := range square {
				backtrack[i] = item
			}

			var success bool
			square, success = insert(tt, square, Coord{x, y})

			if success {

				fillIt(tts, square)

				if len(tts) == 0 {
					os.Exit(0)
				}
			}

			square = backtrack
		}
	}

}

func contains(arr []int, x int) bool {
	for _, item := range arr {
		if x == item {
			return true
		}
	}

	return false
}

func insert(tt TT, input []string, target Coord) ([]string, bool) {
	square := make([]string, len(input))
	for i, s := range input {
		square[i] = s
	}

	coords := lib[tt.ttname]

	for _, c := range coords {
		if target.y+c.y >= len(square) {
			return input, false
		}

		lineCopy := []rune(square[target.y+c.y])

		if target.x+c.x >= len(lineCopy) {
			return input, false
		}

		if lineCopy[target.x+c.x] != '.' {
			return input, false
		}

		lineCopy[target.x+c.x] = rune(tt.name[0])
		square[target.y+c.y] = string(lineCopy)
	}

	return square, true

}

func printSquare(square []string) {

	for _, line := range square {
		fmt.Println(line)
	}

	fmt.Println()
}

func getTetrominoes(tetromino [][]string) []TT {
	tts := []TT{}

	for i, tetro := range tetromino {
		tt := TT{name: string('A' + i)}

		coords := [4]Coord{}

		current := 0
		for y, line := range tetro {
			for x, c := range line {
				if c != '.' {
					coords[current] = Coord{x, y}
					current++
				}
			}
		}

		coords = coordsToZero(coords)

		ttname := getName(coords)

		tt.ttname = ttname

		tts = append(tts, tt)
	}

	return tts
}

func getName(coords [4]Coord) string {

	for name, target := range lib {
		if target == coords {
			return name
		}
	}

	fmt.Println(coords)

	return ""
}

func coordsToZero(coords [4]Coord) [4]Coord {
	xmin, ymin := 4, 4

	for i := 0; i < len(coords); i++ {
		if coords[i].x < xmin {
			xmin = coords[i].x
		}
		if coords[i].y < ymin {
			ymin = coords[i].y
		}
	}

	for i := range coords {
		coords[i].x -= xmin
		coords[i].y -= ymin
	}

	return coords
}
