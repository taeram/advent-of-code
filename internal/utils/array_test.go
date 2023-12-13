package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPriority(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	stringList := []string{"a", "b", "c"}
	a.Equal(true, InArray("a", stringList))
	a.Equal(false, InArray("z", stringList))

	intList := []int{1, 2, 3}
	a.Equal(true, InArray(1, intList))
	a.Equal(false, InArray(4, intList))
}
