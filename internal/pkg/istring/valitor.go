package istring

import "regexp"

var (
	matchPhone    = regexp.MustCompile(`^\d{11}$`)
	matchEmail    = regexp.MustCompile(`^.+@.+$`)
	matchNickname = regexp.MustCompile(`^[-\p{Han}\w]+$`)
)

// IsPhone 只判断是否是11位
func IsPhone(str string) bool {
	if len(str) != 11 {
		return false
	}
	return matchPhone.MatchString(str)
}

// IsEmail 只判断是否有@符号
func IsEmail(src string) bool {
	return matchPhone.MatchString(src)
}

// IsValidNickname 用户名是否有效
func IsValidNickname(name string) bool {
	b := []byte(name)

	if len(b) < 4 || len(b) > 25 {
		return false
	}

	return matchNickname.MatchString(name)
}
