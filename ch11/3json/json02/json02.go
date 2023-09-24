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
	DataOrdered time.Time `json:"date_ordered"` // 注文日時
	CustomerID  string    `json:"customer_id"`  // 顧客ID
	Items       []Item    `json:"items"`        // 商品
}

type Item struct {
	ID   string `json:"id`    // 商品ID
	Name string `json:"name"` // 商品名
}

func main() {
	data := readData()
	var o Order
	err := json.Unmarshal(data, &o)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", o)
}

func readData() []byte {
	b, err := os.ReadFile("../testdata/data.json")
	if err != nil {
		log.Fatal(err)
	}
	return b
}
