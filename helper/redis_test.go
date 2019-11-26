package helper

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	SetupRedis()
	err := RedisSet("tyltr", "nasm", 12)
	fmt.Println(err)
}
