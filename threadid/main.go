//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"fmt"

	"github.com/dollarkillerx/govcl/vcl/rtl"
)

func main() {
	fmt.Println("MainThreadId: ", rtl.MainThreadId())
	fmt.Println("CurrentThreadId: ", rtl.CurrentThreadId())
	go func() {
		fmt.Println("CurrentThreadId2: ", rtl.CurrentThreadId())
	}()
	var s string
	fmt.Scan(&s)
}
