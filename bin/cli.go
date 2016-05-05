package main

import (
	"fmt"
	"github.com/ZhangHang-z/vane"
	"github.com/ZhangHang-z/vane/pkg/cpr"
)

func main() {
	fmt.Println(cpr.ArgI)
	save := vane.DefaultDirIsExist()
	fmt.Println(save)
}
