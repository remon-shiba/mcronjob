package service

import (
	"crypto/rand"
	"fmt"
	"strings"
)

func LineRange() {
	str := "One\nTwo\nThree\nFour\n"

	for part := range strings.Lines(str) {
		print(part)
	}
}

func SplitRange() {
	str := "One-Two-Three-Four"

	for part := range strings.SplitSeq(str, "-") {
		fmt.Println(part)
	}
}

func SplitAfterRange() {
	str := "One-Two-Three-Four"

	for part := range strings.SplitAfterSeq(str, "-") {
		fmt.Println(part)
	}
}

func FieldSeqRange() {
	str := "One Two\nFour Five"

	for part := range strings.FieldsSeq(str) {
		fmt.Println(part)
	}
}
func GenerateRandomText() {
	text := rand.Text()
	fmt.Println(text)
}
