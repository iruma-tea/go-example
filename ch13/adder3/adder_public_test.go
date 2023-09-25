package adder_test

import (
	adder "adder3"
	"testing"
)

func Test_AddNumbers(t *testing.T) {
	result := adder.AddNumbers(2, 3)
	if result != 5 {
		t.Error("結果が正しくない: 想定される結果 5, 得られた結果", result)
	}
}
