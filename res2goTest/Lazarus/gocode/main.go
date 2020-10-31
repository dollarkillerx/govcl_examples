// 由res2go自动生成。
package main

import (
	_ "github.com/dollarkillerx/govcl/pkgs/winappres"
	"github.com/dollarkillerx/govcl/vcl"
)

const Lazarus = true

func main() {

	vcl.Application.Initialize()
	vcl.Application.CreateForm(&MainForm)
	vcl.Application.CreateForm(&About)
	vcl.Application.Run()
}
