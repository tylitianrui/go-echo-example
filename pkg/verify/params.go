package verify

import (
	"fmt"
	"regexp"
)

const (
	MOBILEREGEX = "^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$"
)

var (
	regexMobile *regexp.Regexp
)

func init() {
	regexMobile = regexp.MustCompile(MOBILEREGEX)

}

// 校验手机号是否合法
func VerfifyMobile(m int) bool {
	mobile := fmt.Sprintf("%d", m)
	if len(mobile) != 11 {
		return false
	}
	mobile = regexMobile.FindString(mobile)
	if len(mobile) == 11 {
		return true
	}

	return false

}
