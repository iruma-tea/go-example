package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Order struct {
	ID          string    `json:"id"`           // 注文ID
	DateOrdered time.Time `json:"date_ordered"` // 注文日時
	CustomerID  string    `json:"customer_id"`  // 顧客ID
	Items       []Item    `json:"items"`        // 商品
}

type Item struct {
	ID   string `json:"id"`   // 商品ID
	Name string `json:"name"` // 商品名
}

func main() {
	file, err := os.Open("../testdata/data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var o Order
	d := json.NewDecoder(file)
	d.Decode(&o)
	fmt.Printf("%+v\n", o)
}
