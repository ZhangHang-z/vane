package main

import (
	"fmt"
	"github.com/ZhangHang-z/vane/src/parser"
	"os"
)

/*
git://github.com/bower/bower
git://github.com/bower/bower^1.7.5


https://github.com/bower/bower/archive/master.zip
https://github.com/bower/bower/archive/v1.7.5.zip
*/

// var bowerSRC string = "git://github.com/bower/bower.git"
// var jquerySRC string = "https://code.jquery.com/jquery-2.2.3.min.js"

func main() {
	fmt.Println(os.Args)
	parser.CMDParser()
}
