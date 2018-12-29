package day1

import (
	"fmt"
	"strconv"
	"strings"
)

// Part1 advent challenge
func Part1(changes string) (int, error) {

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
func Part2(changes string) (int, error) {
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
