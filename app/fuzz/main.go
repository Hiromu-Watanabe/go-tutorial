package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
  doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

// // 文字列をバイト単位で反転させる関数
// // => これだとマルチバイトの文字列の際に反転されずにコケる
// func Reverse(s string) string {
// 	b := []byte(s)
// 	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
// 			b[i], b[j] = b[j], b[i]
// 	}
// 	return string(b)
// }

// 文字列をルーン単位で反転させる関数 => マルチバイトの文字列に対応
func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}