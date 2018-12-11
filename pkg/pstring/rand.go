package pstring

import (
	"math/rand"
	"time"
)

const (
	// KC_RAND_KIND_NUM 纯数字
	KC_RAND_KIND_NUM = 0
	// KC_RAND_KIND_LOWER 小写字母
	KC_RAND_KIND_LOWER = 1
	// KC_RAND_KIND_UPPER 大写字母
	KC_RAND_KIND_UPPER = 2
	// KC_RAND_KIND_ALL 数字、大小写字母
	KC_RAND_KIND_ALL = 3
)

// Krand 随机字符串
func Krand(size int, rands ...int) []byte {
	kind := KC_RAND_KIND_ALL
	if len(rands) > 0 {
		kind = rands[0]
	}
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

// RandString 随机字符串
func RandString(size int) string {
	return string(Krand(size, KC_RAND_KIND_ALL))
}

// RandNumberString 随机数字字段串
func RandNumberString(size int) string {
	return string(Krand(size, KC_RAND_KIND_NUM))
}
