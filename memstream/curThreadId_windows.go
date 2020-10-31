package main

import "github.com/dollarkillerx/govcl/vcl/win"

func GetCurrentThreadId() uintptr {
	return win.GetCurrentThreadId()
}
