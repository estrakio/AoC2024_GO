package main

import (
	"AoC2024_GO/day2"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Advent of code 2024")

	//day1.Read_file()
	day2.ReportSafer()
	//test.Test()
	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
