package day2

import (
	"os"
	"strconv"
	"strings"
)

func ReportSafer() {

	filePath := "day2/data.txt" // Chemin relatif depuis le répertoire racine
	data, _ := os.ReadFile(filePath)

	file := strings.Split(string(data), "\r\n")

	soluce := 0
	for _, row := range file {
		checker := 0
		report := strings.Split(string(row), " ")
		for i := 1; i < len(report); i++ {
			if isSafe(report, i, "asc") || isSafe(report, i, "desc") {
				checker++
			}
		}
		if checker == len(report) {
			soluce++
		}
	}

	println(soluce)
}

func isSafe(tab []string, i int, sens string) bool {

	a, _ := strconv.Atoi(tab[i])
	b, _ := strconv.Atoi(tab[i-1])

	diff := abs(a - b)

	if sens == "asc" {
		return a > b && diff <= 3
	} else if sens == "desc" {
		return a < b && diff >= 3
	}

	return false // Invalid sens value
}

// Helper function to calculate absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
