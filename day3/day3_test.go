package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3_Part1(t *testing.T) {
	assert := assert.New(t)
	claims := []string{"#1 @ 1,3: 2x2", "#2 @ 2,3: 2x3", "#3 @ 0,4: 3x1"}
	result, err := part1(claims)
	assert.Nil(err)
	assert.Equal(3, result)
}
func TestDay3_Part2(t *testing.T) {
	assert := assert.New(t)
	claims := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	result, err := part2(claims)
	assert.Nil(err)
	assert.Equal(3, result)
}

func TestDay3_NewClaim(t *testing.T) {
	assert := assert.New(t)
	result, err := newClaim("#1 @ 1,3: 4x5")
	assert.Nil(err)
	assert.Equal(1, result.id)
	assert.Equal(1, result.x0)
	assert.Equal(3, result.y0)
	assert.Equal(4, result.xLen)
	assert.Equal(5, result.yLen)
}

func TestDay3_GetCoverageCoords(t *testing.T) {
	assert := assert.New(t)
	claim, _ := newClaim("#1 @ 1,3: 2x2")
	result := claim.getCoverageCoords()

	assert.EqualValues([]string{"1,3", "1,4", "1,5", "2,3", "2,4", "2,5", "3,3", "3,4", "3,5"}, result)
}

func TestDay3_GetCoverageSquares(t *testing.T) {
	assert := assert.New(t)
	claim, _ := newClaim("#1 @ 1,3: 2x2")
	result := claim.getCoverageSquares()
	assert.EqualValues([]string{"1,3,2,4", "1,4,2,5", "2,3,3,4", "2,4,3,5"}, result)
}

/*
The problem is that many of the claims overlap, causing two or more claims to cover part of the same areas.
For example, consider the following claims:

#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2

Visually, these claim the following areas:

........
...2222.
...2222.
.11XX22.
.11XX22.
.111133.
.111133.
........

The four square inches marked with X are claimed by both 1 and 2.
(Claim 3, while adjacent to the others, does not overlap either of them.)

If the Elves all proceed with their own plans, none of them will have enough fabric.
How many square inches of fabric are within two or more claims?
*/
