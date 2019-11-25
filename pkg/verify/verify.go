package verify

import "go-echo-example/pkg/random"

func GenVerifyCodeString(n int) string {

	return random.GenRandIntN(n)
}

func GenVerifyCodeInt() int {
	return 0
}
