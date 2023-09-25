package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	const data = `
		{"name": "フレッド", "age": 40}
		{"name": "メアリ", "age": 21}
		{"name": "バッド", "age": 30}
	`
	var t struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var r *strings.Reader
	r = strings.NewReader(data)

	var dec *json.Decoder
	dec = json.NewDecoder(r)

	var b bytes.Buffer
	encoder := json.NewEncoder(&b)

	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Printf("t: %v\n", t)
		err = encoder.Encode(t)
		if err != nil {
			panic(err)
		}
	}
	out := b.String()
	fmt.Printf("out:\n%v\n", out)
}
