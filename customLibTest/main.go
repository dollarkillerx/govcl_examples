//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import _ "github.com/dollarkillerx/govcl/samples/customLibTest/alib"
import "github.com/dollarkillerx/govcl/vcl"

type TMainForm struct {
	*vcl.TForm
}

var (
	mainForm *TMainForm
)

func main() {

	vcl.RunApp(&mainForm)
}
