package fn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	is := assert.New(t)
	numbers := []int{1, 2, 3, 4, 5}

	count := 0
	ForEach(numbers, func(n int) {
		count += n
	})
	is.Equal(count, 15)
}

func TestMap(t *testing.T) {
	is := assert.New(t)
	numbers := []int{1, 2, 3, 4, 5}

	var plusOne = func(n int) int {
		return n + 1
	}
	result := Map(numbers, plusOne)
	is.Equal(result, []int{2, 3, 4, 5, 6})
}

func TestFind(t *testing.T) {
	is := assert.New(t)
	numbers := []int{1, 2, 3, 4, 5}

	var equalsTwo = func(n int) bool {
		return n == 2
	}
	is.NotNil(Find(numbers, equalsTwo))

	var equalsTen = func(n int) bool {
		return n == 10
	}
	is.Nil(Find(numbers, equalsTen))
}

func TestFilter(t *testing.T) {
	is := assert.New(t)
	numbers := []int{1, 2, 3, 4, 5}

	var greaterThanTwo = func(i int) bool {
		return i > 2
	}
	is.Equal(Filter(numbers, greaterThanTwo), []int{3, 4, 5})
}

func TestConcat(t *testing.T) {
	is := assert.New(t)
	numbers := []int{1, 2, 3, 4, 5}
	others := []int{6, 7, 8, 9, 10}
	is.Equal(Concat(numbers, others), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func TestEvery(t *testing.T) {
	is := assert.New(t)

	numbers := []int{1, 2, 3, 4, 5}
	var greaterThanZero = func(n int) bool {
		return n > 0
	}
	var greaterThanOne = func(n int) bool {
		return n > 1
	}
	is.True(Every(numbers, greaterThanZero))
	is.False(Every(numbers, greaterThanOne))
}

func TestAny(t *testing.T) {
	is := assert.New(t)

	numbers := []int{1, 2, 3, 4, 5}

	var greaterThanOne = func(n int) bool {
		return n > 1
	}
	var greaterThanFive = func(n int) bool {
		return n > 5
	}
	is.True(Any(numbers, greaterThanOne))
	is.False(Any(numbers, greaterThanFive))
}

func TestIndex(t *testing.T) {
	is := assert.New(t)

	numbers := []int{1, 2, 37, 4, 37}
	is.Equal(Index(numbers, 1), 0)
	is.Equal(Index(numbers, 2), 1)
	is.Equal(Index(numbers, 37), 2)
	is.Equal(Index(numbers, 4), 3)
	is.Equal(Index(numbers, 427), -1)
}

func TestLastIndexOf(t *testing.T) {
	is := assert.New(t)

	numbers := []int{1, 2, 3, 4, 5, 3, 2, 37}
	is.Equal(LastIndexOf(numbers, 1), 0)
	is.Equal(LastIndexOf(numbers, 2), 6)
	is.Equal(LastIndexOf(numbers, 3), 5)
	is.Equal(LastIndexOf(numbers, 237), -1)
}

func TestReduce(t *testing.T) {
	is := assert.New(t)
	numbers := []int{1, 2, 3, 32, 57}

	var add = func(one, two int) int {
		return one + two
	}
	var max = func(one, two int) int {
		if one > two {
			return one
		} else {
			return two
		}
	}

	is.Equal(Reduce(numbers, add), 95)
	is.Equal(Reduce(numbers, max), 57)
	is.Equal(Reduce(numbers, max), Max(numbers))
}

func TestReverse(t *testing.T) {
	is := assert.New(t)
	numbers := []int{1, 2, 3, 4, 5}

	is.Equal(Reverse(numbers), []int{5, 4, 3, 2, 1})
}

func TestPop(t *testing.T) {
	is := assert.New(t)
	numbers := []int{1, 2, 3, 4, 5}

	popped, remaining := Pop(numbers, 3)
	is.Equal(popped, 4)
	is.Equal(remaining, []int{1, 2, 3, 5})

	popped, remaining = Pop(remaining, 2)
	is.Equal(popped, 3)
	is.Equal(remaining, []int{1, 2, 5})

	popped, remaining = Pop(remaining, 0)
	is.Equal(popped, 1)
	is.Equal(remaining, []int{2, 5})

	popped, remaining = Pop(remaining, 1)
	is.Equal(popped, 5)
	is.Equal(remaining, []int{2})

	popped, remaining = Pop(remaining, 0)
	is.Equal(popped, 2)
	is.Equal(remaining, []int{})
}

func TestPush(t *testing.T) {
	is := assert.New(t)

	numbers := []int{1, 2, 3, 4, 5}
	is.Equal(Push(numbers, 27), []int{1, 2, 3, 4, 5, 27})
}

func TestIncludes(t *testing.T) {
	is := assert.New(t)
	numbers := []int{1, 2, 3, 4, 5}

	is.True(Includes(numbers, 1))
	is.False(Includes(numbers, 6))
}

func TestMax(t *testing.T) {
	is := assert.New(t)

	numbers := []int{1, 2, 3, 27, 1, 1, 0, -1}
	is.Equal(Max(numbers), 27)
}

func TestMin(t *testing.T) {
	is := assert.New(t)
	numbers := []int{427, -10, -9, 1, 0, 123545}
	is.Equal(Min(numbers), -10)
}

func TestReplace(t *testing.T) {
	is := assert.New(t)

	numbers := []int{1, 2, 2, 3, 4, 5}
	is.Equal(Replace(numbers, 2, 10), []int{1, 10, 2, 3, 4, 5})
}

func TestReplaceAll(t *testing.T) {
	is := assert.New(t)
	numbers := []int{2, 2, 1, 0, -1, 10, 2, 5}

	is.Equal(ReplaceAll(numbers, 2, 35), []int{35, 35, 1, 0, -1, 10, 35, 5})
}

func TestDelete(t *testing.T) {
	is := assert.New(t)
	numbers := []int{1, 2, 3, 4, 5, 4}

	is.Equal(Delete(numbers, 4), []int{1, 2, 3, 5, 4})
}

func TestDeleteAll(t *testing.T) {
	is := assert.New(t)
	numbers := []int{1, 2, 3, 4, 1, 37, 0, -1, 10, 99, 1}

	is.Equal(DeleteAll(numbers, 1), []int{2, 3, 4, 37, 0, -1, 10, 99})
}