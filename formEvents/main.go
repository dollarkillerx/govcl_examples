package main

import (
	_ "github.com/dollarkillerx/govcl/pkgs/macapp"
	_ "github.com/dollarkillerx/govcl/pkgs/winappres"
	"github.com/dollarkillerx/govcl/vcl"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&Form1)
	vcl.Application.CreateForm(&Form2)

	vcl.Application.Run()

}
