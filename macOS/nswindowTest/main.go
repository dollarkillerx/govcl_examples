// +build darwin

package main

import (
	_ "github.com/dollarkillerx/govcl/pkgs/winappres"
	"github.com/dollarkillerx/govcl/vcl"
)

func main() {
	vcl.RunApp(&Form1)
}
