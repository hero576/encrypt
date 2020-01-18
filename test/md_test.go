package test

import (
	"encrypt"
	"fmt"
	"testing"
)

func Test_Md(t *testing.T) {
	if res, err := encrypt.HASH("123456", "md5", false); err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}
func Test_MdDouble(t *testing.T) {
	if res, err := encrypt.HASHDouble("123456", "md5", false); err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}
