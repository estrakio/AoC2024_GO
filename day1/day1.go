package day1

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func Read_file() {

	filePath := "day1/data.txt" // Chemin relatif depuis le répertoire racine
	data, _ := os.ReadFile(filePath)

	result := strings.Split(string(data), "\r\n")
	//fmt.Printf("%#v", result)

	left := []int{}
	right := []int{}

	var leftData int
	var rightData int

	for _, row := range result {

		fmt.Sscanf(row, "%d %d", &leftData, &rightData)
		left = append(left, leftData)
		right = append(right, rightData)

	}

	sort.Ints(left)
	sort.Ints(right)
	partOne(left, right)
	partTwo(left, right)
}

func partTwo(left []int, right []int) {
	dico := make(map[int]int)
	result := 0

	for i := range right {
		dico[right[i]]++
	}

	//fmt.Printf("%#v\n", dico)
	for _, val := range left {
		result += val * dico[val]
	}

	println("solution partie 2 : ", result)

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func partOne(left []int, right []int) {

	soluce := 0

	for i := range left {
		soluce += Abs(left[i] - right[i])
	}
	println("solution partie 1 : ", soluce)
}

//fmt.Println("\n right")
//fmt.Printf("%#v\n", right)
//fmt.Println("\n left")
//fmt.Printf("%#v\n", left)
