package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string { // 値レシーバ（cのコピーが渡される）
	return fmt.Sprintf("合計: %d, 更新: %v", c.total, c.lastUpdated)
} //listend1

func main() {
	var c Counter
	fmt.Println(c)
	(&c).Increment()
	fmt.Println(c)
}
