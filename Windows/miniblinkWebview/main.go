package main

import (
	"github.com/dollarkillerx/govcl/pkgs/miniblink"
	_ "github.com/dollarkillerx/govcl/pkgs/winappres"
	"github.com/dollarkillerx/govcl/vcl"
)

func main() {

	miniblink.Init()
	defer miniblink.Finalize()

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&MainForm)
	vcl.Application.CreateForm(&HTMLForm)
	vcl.Application.CreateForm(&JsForm)
	vcl.Application.Run()

}
