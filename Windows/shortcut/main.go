// +build windows

package main

import (
	"os"

	_ "github.com/dollarkillerx/govcl/pkgs/winappres"
	"github.com/dollarkillerx/govcl/vcl"
	"github.com/dollarkillerx/govcl/vcl/rtl"
	"github.com/dollarkillerx/govcl/vcl/win"
)

func main() {
	rtl.CreateURLShortCut(win.GetDesktopPath(), "govcl", "https://github.com/dollarkillerx/govcl")

	rtl.CreateShortCut(win.GetDesktopPath(), "shortcut", os.Args[0], "", "描述", "-b -c")

	vcl.ShowMessage("Hello!")

}
