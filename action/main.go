package main

import "github.com/dollarkillerx/govcl/vcl"
import _ "github.com/dollarkillerx/govcl/pkgs/winappres"

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&mainForm, true)
	vcl.Application.Run()
}
