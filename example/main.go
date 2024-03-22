package main

import (
	"fmt"

	"github.com/crowdsecurity/go-tokenizers"
)

func main() {
	tk, err := tokenizers.FromFile("./tokenizer.json")
	if err != nil {
		panic(err)
	}
	// release native resources
	defer tk.Close()
	fmt.Println("Vocab size:", tk.VocabSize())
	// Vocab size: 30522
	fmt.Println(tk.Encode("hello world", false))
	// [2829 4419 14523 2058 1996 13971 3899] [hello world]
	fmt.Println(tk.Encode("hello world", true))
	// [101 2829 4419 14523 2058 1996 13971 3899 102] [[CLS] hello world [SEP]]
	fmt.Println(tk.Decode([]uint32{2829, 4419, 14523, 2058, 1996, 13971, 3899}, true))
	// hello world
}
