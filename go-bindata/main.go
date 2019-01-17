package main

import "strings"

var words map[string]string

func init() {
	bin, _ := Asset("hello") // アセットを読み込み
	c := string(bin)
	w := strings.Split(c, "=")
	words[w[0]] = w[1]
}

func Message(key string) string {
	return words[key]
}
