package random

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 随机生成n位数字字符串
func GenRandIntN(n int) string {
	return fmt.Sprintf(
		"%0"+strconv.Itoa(n)+"v",
		rand.New(rand.NewSource(
			time.Now().UnixNano())).Int31n(int32(Pow(10, n))))
}

// x 的n次方
func Pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	result := calPow(x, n)
	if n < 0 {
		result = 1 / result
	}
	return result
}

func calPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	// 向右移动一位
	result := calPow(x, n>>1)
	result *= result

	// 如果n是奇数
	if n&1 == 1 {
		result *= x
	}

	return result
}
