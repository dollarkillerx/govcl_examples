// 在这里写你的事件

package main

import "github.com/dollarkillerx/govcl/vcl"

//::private::
type THTMLFormFields struct {
}

func (f *THTMLForm) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
}
