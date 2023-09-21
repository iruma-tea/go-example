package foo

import "internal_example/foo/internal"

func UseDoubler(i int) int {
	return internal.Doubler(i)
}
