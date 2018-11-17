package main

import (
	"regexp"
	"strings"
)

var (
	datasetIdExp = regexp.MustCompile("/[a-zA-Z0-9]{5}/aaaaa/([a-zA-Z0-9]{5})")
	hoge         = "/55555/aaaaa/11111"

	sourceValidator = func(w string, matcher []string) bool {
		for _, v := range matcher {
			if v == w {
				return true
			}
		}
		return false
	}
)

func main() {
}

func Strings() {
	strings.Split(hoge, "/")

}

func RegExp() {
	datasetIdExp.FindStringSubmatch(hoge)
}

func Matcher() {
	sourceValidator("hoge", []string{hoge})
}

func FuncMatcher() {
	Validate("hoge", []string{hoge})
}

func Validate(w string, matcher []string) bool {
	for _, v := range matcher {
		if v == w {
			return true
		}
	}
	return false
}
