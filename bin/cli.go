package main

import (
	"fmt"
	"github.com/ZhangHang-z/vane/src/vane"
	//"github.com/ZhangHang-z/vane/src/cmd"
	"github.com/ZhangHang-z/vane/src/down"
)

var bowerSRC string = "git://github.com/bower/bower.git"
var jquerySRC string = "https://code.jquery.com/jquery-2.2.3.min.js"

func main() {
	save := vane.DftDirIsExist()
	fmt.Println(save)

	err := vane.MkDftDir()
	fmt.Println("Create default install directory: ", err)

	save = vane.DftDirIsExist()
	fmt.Println(save)

	down.GotoComponentsDir("vane_components")
	buf := down.GitRepDownloader(bowerSRC)
	fmt.Println(buf.String())
}
