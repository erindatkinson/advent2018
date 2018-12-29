package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2_CheckTwoMatch(t *testing.T) {
	assert := assert.New(t)
	ids := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	twoMatch := []bool{false, true, true, false, true, true, false}
	for index := 0; index < len(ids); index++ {
		result := checkTwoMatch(ids[index])
		assert.Equal(twoMatch[index], result)
	}
}

func TestDay2_CheckThreeMatch(t *testing.T) {
	assert := assert.New(t)
	ids := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	threeMatch := []bool{false, true, false, true, false, false, true}

	for index := 0; index < len(ids); index++ {
		result := checkThreeMatch(ids[index])
		assert.Equal(threeMatch[index], result)
	}
}

func TestDay2_Part1(t *testing.T) {
	//assert := assert.New(t)

}

func TestDay2_Checksum(t *testing.T) {
	assert := assert.New(t)
	result := checksum(3, 4)
	assert.Equal(12, result)

}

/*

    abcdef contains no letters that appear exactly two or three times.
    bababc contains two a and three b, so it counts for both.
    abbcde contains two b, but no letter appears exactly three times.
    abcccd contains three c, but no letter appears exactly two times.
    aabcdd contains two a and two d, but it only counts once.
    abcdee contains two e.
    ababab contains three a and three b, but it only counts once.


		Of these box IDs, four of them contain a letter which
		appears exactly twice, and three of them contain a letter
		which appears exactly three times. Multiplying these
		together produces a checksum of 4 * 3 = 12.
*/

func TestDay2_CheckDifferenceCount(t *testing.T) {
	assert := assert.New(t)
	test := "abcde"
	ids := []string{"fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	results := []int{5, 5, 5, 5, 2, 5}
	for index := 0; index < len(ids); index++ {
		result, err := checkDifferenceCount(test, ids[index])
		assert.Nil(err)
		assert.Equal(results[index], result)
	}
}

func TestDay2_FilterCommon(t *testing.T) {
	assert := assert.New(t)
	test := "abcde"
	id := "axcye"
	result := filterCommon(test, id)
	assert.Equal("ace", result)

}

func TestDay2_Part2(t *testing.T) {
	assert := assert.New(t)
	ids := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	result, err := part2(ids)
	assert.Nil(err)
	assert.Equal("fgij", result)
}

/*
abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz

The IDs abcde and axcye are close, but they differ by two
characters (the second and fourth). However, the IDs fghij
and fguij differ by exactly one character, the third
(h and u). Those must be the correct boxes.

What letters are common between the two correct box IDs?
(In the example above, this is found by removing the
	differing character from either ID, producing fgij.)
*/
