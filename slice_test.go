package slicex

import . "github.com/pkg4go/assert"
import "testing"

func TestContains(t *testing.T) {
	a := A{t}

	a.Equal(Contains([]int{1, 2}, 1), true)
	a.Equal(Contains([]int{1, 2}, 3), false)
	a.Equal(Contains([]int{1, 2}, "1"), false)

	a.Equal(Contains([]string{"1", "2"}, "1"), true)
	a.Equal(Contains([]string{"1", "2"}, "3"), false)
	a.Equal(Contains([]string{"1", "2"}, 1), false)
}

func TestIndex(t *testing.T) {
	a := A{t}

	a.Equal(Index([]int{1, 2}, 1), 0)
	a.Equal(Index([]int{1, 2}, 2), 1)
	a.Equal(Index([]int{1, 2}, 3), -1)
	a.Equal(Index([]int{1, 2}, "1"), -1)

	a.Equal(Index([]string{"1", "2"}, "1"), 0)
	a.Equal(Index([]string{"1", "2"}, "3"), -1)
	a.Equal(Index([]string{"1", "2"}, 1), -1)
}
