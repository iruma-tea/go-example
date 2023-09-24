package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	dataToFile := Person{
		Name: "フレッド",
		Age:  40,
	}

	tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")

	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())
	err = json.NewEncoder(tmpFile).Encode(dataToFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ファイルに書き込んだデータ: %+v\n", dataToFile) // %vで構造体をフィールド名付きで表示
	// 読み込み
	tmpFile2, err := os.Open(tmpFile.Name()) //liststart3
	if err != nil {
		panic(err)
	}
	var dataFromFile Person
	err = json.NewDecoder(tmpFile2).Decode(&dataFromFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ファイルから読み込んだデータ: %+v\n", dataFromFile) // 構造体をフィールド名付きで表示

}
