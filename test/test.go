package test

import (
	"os"
	"strconv"
	"strings"
)

func Test() {

	filePath := "day2/data.txt" // Chemin relatif depuis le répertoire racine
	data, _ := os.ReadFile(filePath)

	file := strings.Split(string(data), "\r\n")
	soluce := 0
	for _, row := range file {
		soluce += rowIsSafe(strings.Split(string(row), " "))
	}
	println(soluce)
}

func rowIsSafe(row []string) int {
	compteurAsc := 0
	compteurDesc := 0
	arrSize := len(row) - 1
	for i := 1; i < len(row); i++ {
		first, _ := strconv.Atoi(row[i])
		second, _ := strconv.Atoi(row[i-1])
		if isThree(first - second) {
			if isFollowing(first, second, "asc") {
				compteurAsc++
			}
			if isFollowing(first, second, "desc") {
				compteurDesc++
			}
		}
	}
	if compteurAsc == arrSize || compteurDesc == arrSize {
		return 1
	}
	return 0
}

func isFollowing(a int, b int, sens string) bool {
	if sens == "asc" {
		if a > b {
			return true
		}
	} else {
		if a < b {
			return true
		}
	}
	return false
}

func isThree(x int) bool {
	test := abs(x)
	if test <= 3 {
		return true
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
