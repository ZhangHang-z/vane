package main

import (
	"flag"
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

	parser.CMDParser()
	fmt.Println(os.Args)
	parser.OPTParser()
	fmt.Println(parser.ArgI, parser.ArgInstall, parser.ArgIsOpt, parser.ArgPkgs, parser.Save)
	fmt.Println(flag.Args())
	fmt.Println(flag.Arg(0))
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())

}
