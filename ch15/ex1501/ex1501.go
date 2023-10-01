package main

import (
	"fmt"
	"strings"
)

type Orderable interface {
	//  メソッドOrderは次のような値を戻す：
	//  Orderableが渡された値よりも小さい場合、0よりも小さい値を返す
	//  Orderableが渡された値よりも大きい場合、0よりも大きい値を返す
	//  2つの値が等しい時0を返す
	Order(interface{}) int
}

type Tree struct {
	val         Orderable
	left, right *Tree
}

func (t *Tree) Insert(val Orderable) *Tree {
	if t == nil {
		return &Tree{val: val}
	}

	switch comp := val.Order(t.val); {
	case comp < 0:
		t.left = t.left.Insert(val)
	case comp > 0:
		t.right = t.right.Insert(val)
	}
	return t
}

type OrderableInt int

func (oi OrderableInt) Order(val interface{}) int {
	return int(oi - val.(OrderableInt))
}

type OrderableString string

func (os OrderableString) Order(val interface{}) int {
	return strings.Compare(string(os), val.(string))
}

func main() {
	var it *Tree
	it = it.Insert(OrderableInt(5))
	it = it.Insert(OrderableInt(3))
	fmt.Println(it.val)

	{
		var it *Tree
		it = it.Insert(OrderableInt(5))
		it = it.Insert(OrderableString("nope")) // 実行時のエラー
	}
}
