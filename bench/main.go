package main

import (
	"regexp"
	"strings"
)

var (
	datasetIdExp = regexp.MustCompile("/[a-zA-Z0-9]{5}/aaaaa/([a-zA-Z0-9]{5})")
	hoge         = "/55555/aaaaa/11111"
)

func main() {
}

func Strings() {
	strings.Split(hoge, "/")

}

func RegExp() {
	datasetIdExp.FindStringSubmatch(hoge)
}
