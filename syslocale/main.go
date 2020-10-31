package main

import (
	"fmt"

	_ "github.com/dollarkillerx/govcl/pkgs/winappres"
	"github.com/dollarkillerx/govcl/vcl/rtl"
)

func main() {
	fmt.Println(rtl.SysLocale)
}
