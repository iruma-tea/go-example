package main

import (
	"fmt"
	"strings"
	"sync"
)

type dictionary map[string]string
type translator struct {
	dict dictionary
}

func (t translator) translateOneLang(s string, ch chan<- string) {
	words := strings.Split(s, " ")
	var r string
	for _, w := range words {
		if t.dict[w] != "" {
			r += t.dict[w] + " "
		} else {
			r += w + " "
		}
	}
	ch <- r
}

func (t translator) initTranslator(d dictionary) translator {
	t.dict = d
	return t
}

// ファイルから読み込んで作るのが一般的
var ejDict = dictionary{ // 英->日 辞書
	"I":         "私",
	"We":        "我々",
	"write":     "書く",
	"translate": "翻訳する",
	"see":       "見る",
	"books":     "本",
	"mountains": "山",
	"river":     "川",
	"library":   "図書館",
}

var gjDict = dictionary{ // 独->日 辞書
	"Ich":       "私",
	"Wir":       "我々",
	"schreibe":  "書く",
	"übersetze": "翻訳する",
	"sehen":     "見る",
	"Bücher":    "本",
	// "Berge": "山",
	"Fluss":      "川",
	"Bibliothek": "図書館",
}

var once sync.Once // 一度だけ実行するため。本文参照

func main() {
	fmt.Println(Translate("I translate books"))
	fmt.Println(Translate("I write books"))
	fmt.Println(Translate("We see mountains"))
	fmt.Println(Translate("Ich schreibe Bücher"))
	fmt.Println(Translate("Ich übersetze Bücher"))
	fmt.Println(Translate("Wir sehen Berge")) // 未登録語あり
}

var translators []translator

func Translate(origin string) []string {
	fmt.Println("原文:", origin)
	var t translator
	once.Do(func() { // 最初に呼ばれた時だけ辞書を読み込み
		fmt.Println("*初期化処理を実行* ")
		translators = append(translators, t.initTranslator(ejDict))
		translators = append(translators, t.initTranslator(gjDict))
	})
	ch := make(chan string) // 結果送受信用のチャネル
	for _, t := range translators {
		// 各言語についてトランスレータを並行実行
		go t.translateOneLang(origin, ch)
	}

	var r []string
	for i := 0; i < len(translators); i++ {
		// トランスレータの数だけ翻訳結果をもらう
		translation := <-ch
		translation = strings.TrimSpace(translation) //余分な空白とる
		// fmt.Println("結果:", translation)
		if origin != translation { // 原文と違う結果ができたら
			r = append(r, translation) // 結果を記憶
		}
	}
	return r
}
