package sibling

import "internal_example/foo/internal"

func AlsoUseDoubler(i int) int {
	return internal.Doubler(i)
}
