package chapter3

import (
	"fmt"
	"unicode/utf8"
)

func stringFunc() {
	s := "阿斯蒂芬yes"

	// 15，s是utf8编码，一个汉字3字节
	fmt.Print(len(s))

	for i, ch := range s {
		// i是第几个字节，ch是utf8编码号
		fmt.Print(i, ch)
	}

	count := utf8.RuneCountInString(s)
	fmt.Print(count) //7

	for i, ch := range []rune(s) {
		// i是第几个字符，ch是字符
		fmt.Print(i, ch)
	}

	//strings包有很多常用字符串处理方法

}
