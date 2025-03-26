package pwd

import (
	"fmt"
	"testing"
)

var hash string

func TestHashPwd(t *testing.T) {
	hash = HashPwd("zyuuforyu")
	fmt.Println(hash)
}

func TestVerifyPwd(t *testing.T) {
	VerifyPwd("123456", hash)
}
