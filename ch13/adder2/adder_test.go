package adder

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("テストの準備")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("テスト後の後片付け")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("TestFirstではTestMainで設定されたものを使う", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("TestSecondでもTestMainで設定されたものを使う", testTime)
}

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("結果が違う: 期待する結果 5, 得られた結果", result)
	}
}
