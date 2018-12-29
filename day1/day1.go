package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/erindatkinson/advent2018/utils"
)

// Execute entrypoint into day1
func Execute(filename string) {
	fmt.Println("=== Day1 ===")
	day1file, err := utils.ReadFile("files/day1.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result, err := part1(day1file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Part 1 Result: %d\n", result)

	result, err = part2(day1file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 Result: %d\n", result)
}

// Part1 advent challenge
func part1(changes string) (int, error) {

	calibration := 0
	lines := strings.Split(changes, "\n")
	for i := 0; i < len(lines); i++ {
		operation, value, err := getChangeInteger(lines[i])
		if err != nil {
			return 0, err
		}

		switch operation {
		case "+":
			calibration = calibration + value
		case "-":
			calibration = calibration - value
		}
	}

	return calibration, nil
}

// Part2 advent Challege
func part2(changes string) (int, error) {
	// for each calibration, check if it's been seen before, first to match, return
	calibration := 0
	var checks map[int]bool
	checks = make(map[int]bool)
	checks[calibration] = true
	lines := strings.Split(changes, "\n")
	for duplicate := false; ; {
		for i := 0; i < len(lines); i++ {
			operation, value, err := getChangeInteger(lines[i])
			if err != nil {
				return 0, err
			}
			switch operation {
			case "+":
				calibration = calibration + value
			case "-":
				calibration = calibration - value
			}

			_, duplicate = checks[calibration]
			if duplicate {
				return calibration, nil
			}
			checks[calibration] = true
		}
	}
}

func getChangeInteger(change string) (string, int, error) {
	value, err := strconv.Atoi(change[1:])
	if err != nil {
		return "", 0, fmt.Errorf("Error converting value to int: %s", err)
	}
	return string(change[0]), value, nil
}
