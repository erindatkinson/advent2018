package main

import (
	"fmt"
	"os"

	"github.com/erindatkinson/advent2018/day1"
	"github.com/erindatkinson/advent2018/utils"
)

func main() {
	fmt.Println("=== Day1 ===")
	day1file, err := utils.ReadFile("files/day1.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	d1p1(day1file)
	d1p2(day1file)

}

func d1p1(read string) {

	result, err := day1.Part1(read)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 Result: %d\n", result)
}

func d1p2(read string) {
	result, err := day1.Part2(read)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 Result: %d\n", result)
}
