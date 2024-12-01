package main

import (
	"AoC2024_GO/day1"
	"fmt"
	"time"
)

func first_day() {
	day1.Read_file()
}

func main() {
	start := time.Now()
	fmt.Println("Ewan il est pd")
	first_day()
	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
