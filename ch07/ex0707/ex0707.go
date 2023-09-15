package main

import "fmt"

func main() {
	fmt.Println("===== 7.2.7　iotaと列挙型 =====")
	{
		type MailCategory int

		const (
			Uncategorized MailCategory = iota
			Personal
			Spam
			Social
			Advertisements
		)

		m := Personal
		fmt.Println("Personal:", m)
		m = Uncategorized
		fmt.Println("Uncategorized:", m)
	}
}
