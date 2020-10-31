package main

import (
	_ "github.com/dollarkillerx/govcl/pkgs/winappres"
	"github.com/dollarkillerx/govcl/vcl"
	"github.com/dollarkillerx/govcl/vcl/i18n"
)

func main() {

	i18n.InitDefaultLang()
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&Form1)
	vcl.Application.Run()

}
