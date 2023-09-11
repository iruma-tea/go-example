package main

import "fmt"

func main() {
	fmt.Println("===== 2.1　基本型 =====")
	fmt.Println("===== 2.1.1 ゼロ値 =====")
	{
		var x int
		var y float32
		var z string

		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Printf("z:|%s|\n", z)
	}

	fmt.Println("===== 2.1.2　リテラル =====")
	fmt.Println("2.1.2.1　整数リテラル")
	fmt.Println(1_234_567) // 1234567

	fmt.Println("2.1.2.2　浮動小数点数リテラル")
	fmt.Println(1_234.567_89) // 1234.56789

	fmt.Println("2.1.2.3　runeリテラル")
	{
		// A
		fmt.Println("1->\x41")
		fmt.Println("2->\u0041")
		fmt.Println("3->\U00000041")

		// あ
		fmt.Println("1->\u3042")
		fmt.Println("2->\U00003042")

		// 絵文字
		fmt.Println("1->\U0001F600")

		fmt.Println("\n2.1.2.4　文字列リテラル")
		x := `バッククオートを使って"ロー文字列リテラル"を書くことで
改行や二重引用符（ダブルクオート）を文字列中に含めることができる`
		fmt.Println(x)
	}

	fmt.Println("\n2.1.2.5　リテラルと型")
	{
		var x float32 = 2 * 0.23 * 3
		fmt.Println(x)
		var a float64 = 3.14
		var b float64 = 10 / a
		fmt.Println(b)

		// var c int = "abc"  // コンパイル時のエラーとなる
		// var s string = 12  // コンパイル時のエラーとなる
		var d int
		// d = 12.3 // コンパイル時のエラー(桁落ちする)
		d = 12.0
		fmt.Println("d:", d)
	}

	fmt.Println("===== 2.1.3 論理型 =====")
	{
		var flag bool
		var isAwesome = true
		fmt.Println(flag)
		fmt.Println(isAwesome)
	}

	fmt.Println("===== あふれる =====")
	{
		// var x byte = 1000 // エラー。たとえば1000をbyte型の変数に代入しようとするとエラーになります。
		var x byte = 100
		fmt.Println(x)
	}

	fmt.Println("===== 2.1.4.4　整数関連の演算子 =====")
	{
		var x int = 10
		x *= 2
		fmt.Println(x)
	}

	fmt.Println("===== 2.1.4.5　浮動小数点数型 =====")
	{
		var x float64 = 0
		fmt.Println(x) // 0
		x += 1.333456
		fmt.Println(x) // 1.333456

		x = 1.5
		var y float64 = 2.5
		x /= y
		fmt.Println(x) // 0.6
		// 		x = x % y    // エラー。 floatにたいして「%」は使えない

		fmt.Println("x/0=", x/0)   // +Inf
		fmt.Println("-x/0=", -x/0) // -Inf
	}

}
