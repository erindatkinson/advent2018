package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/erindatkinson/advent2018/utils"
)

// Execute Day3
func Execute(filename string) {
	fmt.Println("=== Day3 ===")
	day3file, err := utils.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	claims := strings.Split(day3file, "\n")
	result1, err := part1(claims)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Part1 Result: %d\n", result1)

	result2, err := part2(claims)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Part1 Result: %d\n", result2)

}

func part1(data []string) (int, error) {
	overlap := 0
	var claimSquares [][]string

	for _, value := range data {
		claimObj, err := newClaim(value)
		if err != nil {
			return -1, err
		}
		claimSquares = append(claimSquares, claimObj.getCoverageSquares())
	}

	overlapMap := getOverlapMap(claimSquares)

	for _, coverCount := range overlapMap {
		if coverCount > 1 {
			overlap++
		}
	}
	return overlap, nil
}

func part2(data []string) (int, error) {
	var claimSquares [][]string
	var claimObjs []Claim

	for _, value := range data {
		claimObj, err := newClaim(value)
		if err != nil {
			return -1, err
		}
		claimObjs = append(claimObjs, claimObj)
		claimSquares = append(claimSquares, claimObj.getCoverageSquares())
	}

	overlapMap := getOverlapMap(claimSquares)

	for claimIndex, squares := range claimSquares {
		fmt.Printf(". ")
		badClaim := false
		for _, square := range squares {
			if !badClaim {
				for overlapSquare, count := range overlapMap {
					if !badClaim {
						if count > 1 {
							if strings.Compare(overlapSquare, square) == 0 {
								badClaim = true
							}
						}
					}
				}
			}
		}
		if !badClaim {
			fmt.Println("")
			return claimObjs[claimIndex].id, nil
		}
	}
	return -1, fmt.Errorf("No non-overlapping claims present")
}

func getOverlapMap(claimSquares [][]string) map[string]int {
	cover := make(map[string]int)
	for _, val := range claimSquares {
		for _, square := range val {
			_, present := cover[square]
			if !present {
				cover[square] = 1
			} else {
				cover[square]++
			}
		}
	}
	return cover
}

// Claim struct to manage data
type Claim struct {
	id   int
	x0   int
	y0   int
	xLen int
	yLen int
}

func (c Claim) getCoverageCoords() []string {
	var coverageCoords []string
	for x, xCount := c.x0, 0; xCount <= c.xLen; xCount++ {
		for y, yCount := c.y0, 0; yCount <= c.yLen; yCount++ {
			coverageCoords = append(coverageCoords, fmt.Sprintf("%d,%d", x+xCount, y+yCount))
		}
	}
	return coverageCoords
}
func (c Claim) getCoverageSquares() []string {
	coverageCoords := c.getCoverageCoords()
	var coverageSquares []string
	for index, value := range coverageCoords {
		valueInts := strings.Split(value, ",")

		testSquareValue := fmt.Sprintf(
			"%d,%d",
			mustConvertAtoi([]byte(valueInts[0]))+1,
			mustConvertAtoi([]byte(valueInts[1]))+1)

		for subIndex := index + 1; subIndex < len(coverageCoords); subIndex++ {
			if strings.Compare(coverageCoords[subIndex], testSquareValue) == 0 {
				coverageSquares = append(coverageSquares, fmt.Sprintf("%s,%s", value, testSquareValue))
			}
		}
	}
	return coverageSquares
}

func newClaim(data string) (Claim, error) {
	pattern, err := regexp.Compile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")
	if err != nil {
		return Claim{}, err
	}
	var dataInt [5]int
	dataBytes := []byte(data)
	allIndices := pattern.FindAllSubmatchIndex(dataBytes, -1)
	for _, location := range allIndices {
		for index := 1; index < len(location)/2; index++ {
			dataInt[index-1] = mustConvertAtoi(dataBytes[location[2*index]:location[2*index+1]])
		}
	}
	return Claim{id: dataInt[0], x0: dataInt[1], y0: dataInt[2], xLen: dataInt[3], yLen: dataInt[4]}, nil
}

func mustConvertAtoi(data []byte) int {
	result, err := strconv.Atoi(string(data))
	if err != nil {
		fmt.Printf("Error parsing integer (%s) in mustConvertAtoi function", data)
		os.Exit(1)
	}
	return result
}
