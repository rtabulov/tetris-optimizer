package main

import (
	"fmt"
)

func main() {
	// data, err := ioutil.ReadFile("test.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// arr := parseData(data)

	item := "#...\n#...\n#...\n#...\n"
	result := "....\n....\n..##\n..##\n"

	// for _, item := range arr {
	// 	result = combine(result, item)

	// 	fmt.Println(result)
	// }

	result = insert(result, item)
	fmt.Print(result)

	// fmt.Println(arr)
}

func combine(result, item string) string {
	result = squarify(result)

	return insert(result, item)
}

func insert(result, item string) string {
	// for {
	// 	if fits(result, item) {
	// 		break
	// 	}

	// result = indcreaseSize(result)
	// }

	return result + item
}

func fits(result, item string) bool {

	return false
}

func indcreaseSize(str string) string {

	w := getWidth(str)

	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			str = str[:i] + "." + str[i:]
			i++
		}
	}

	for i := 0; i <= w; i++ {
		str += "."
	}

	return str + "\n"
}

func squarify(str string) string {
	h := getHeight(str)
	w := getWidth(str)

	fmt.Println(h, w)

	return ""
}

func getWidth(str string) int {
	result := 0
	for _, c := range str {
		if c == '\n' {
			break
		}
		result++
	}

	return result
}

func getHeight(str string) int {
	result := 0
	for _, c := range str {
		if c == '\n' {
			result++
		}
	}

	return result
}

func parseData(data []byte) []string {
	len := countTetrominoes(data)
	result := make([]string, len)
	lines := 0
	curIndex := 0
	for _, c := range data {
		// fmt.Println(curIndex)
		if c == '\n' {
			lines++
		}

		if lines == 5 {
			lines = 0
			curIndex++
			continue
		}

		result[curIndex] += string(c)
	}

	return result
}

func countTetrominoes(data []byte) int {
	len := 2
	for _, c := range data {
		if c == '\n' {
			len++
		}
	}

	// fmt.Println(len)

	return len / 5
}
