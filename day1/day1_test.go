package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1_Part1(t *testing.T) {
	assert := assert.New(t)
	example1 := "+1\n+1\n+1"
	result, err := Part1(example1)
	assert.Nil(err)
	assert.Equal(3, result)

	example2 := "+1\n+1\n-2"
	result, err = Part1(example2)
	assert.Nil(err)
	assert.Equal(0, result)

	example3 := "-1\n-2\n-3"
	result, err = Part1(example3)
	assert.Nil(err)
	assert.Equal(-6, result)

}

func TestDay1_Part2(t *testing.T) {
	assert := assert.New(t)
	example1 := "+1\n-1"
	result, err := Part2(example1)
	assert.Nil(err)
	assert.Equal(0, result)

	example2 := "+3\n+3\n+4\n-2\n-4"
	result, err = Part2(example2)
	assert.Nil(err)
	assert.Equal(10, result)

	example3 := "-6\n+3\n+8\n+5\n-6"
	result, err = Part2(example3)
	assert.Nil(err)
	assert.Equal(5, result)

	example4 := "+7\n+7\n-2\n-7\n-4"
	result, err = Part2(example4)
	assert.Nil(err)
	assert.Equal(14, result)
}
