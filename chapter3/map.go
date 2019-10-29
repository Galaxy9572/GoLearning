package chapter3

import "fmt"

// map的key使用哈希表，必须可以比较相等
// 除了slice，map，function的其他类型都可以当做key
// Struct类型不包含以上类型也可以作为key
func mapFunc() {

	// map定义方式1：
	// map[key]value{key:value,key:value}
	m := map[string]string{
		"go": "golang",
		"c":  "clang",
	}

	fmt.Print(m)

	for k, v := range m {
		fmt.Print(k, v)
	}

	// map定义方式2：
	// make(map[key]value)
	m2 := make(map[string]int) // m2 = nil
	fmt.Print(m2)

	var m3 map[string]int // empty map

	fmt.Print(m3)

	// 获取元素，检查元素是否存在
	golang, exists := m["go"]
	fmt.Print(golang, exists) // golang,true

	// 获取元素，检查元素是否存在
	cpp, exists := m["cpp"]
	fmt.Print(cpp, exists) //空字符串,false，golang的map取不到值不是返回nil，而是value对应类型的空值

	// 删除map中的元素
	delete(m, "c")
}
