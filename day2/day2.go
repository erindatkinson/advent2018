package day2

import (
	"fmt"
	"os"
	"strings"

	"github.com/erindatkinson/advent2018/utils"
)

// Execute entrypoint into day2
func Execute(filename string) {
	fmt.Println("=== Day2 ===")
	day2file, err := utils.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ids := strings.Split(day2file, "\n")
	result := part1(ids)
	fmt.Printf("Part 1 Result: %d\n", result)

}

func part1(ids []string) int {
	twoMatchCount, threeMatchCount := 0, 0
	for index := 0; index < len(ids); index++ {
		twoMatch := checkTwoMatch(ids[index])
		if twoMatch {
			twoMatchCount = twoMatchCount + 1
		}
		threeMatch := checkThreeMatch(ids[index])

		if threeMatch {
			threeMatchCount = threeMatchCount + 1
		}
	}
	return checksum(twoMatchCount, threeMatchCount)

}

func checkTwoMatch(id string) bool {
	return checkMatch(2, id)
}

func checkThreeMatch(id string) bool {
	return checkMatch(3, id)
}

func checksum(twoMatch int, threeMatch int) int {
	return twoMatch * threeMatch
}

func checkMatch(count int, id string) bool {
	var checks map[string]int
	checks = make(map[string]int)
	var exists bool
	for index := 0; index < len(id); index++ {
		char := string(id[index])
		_, exists = checks[char]
		if exists {
			checks[char] = checks[char] + 1
		} else {
			checks[char] = 1
		}
	}
	for _, v := range checks {
		if v == count {
			return true
		}
	}
	return false
}
