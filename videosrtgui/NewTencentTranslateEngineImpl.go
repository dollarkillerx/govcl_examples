// Automatically generated by the simple GoVCL IDE.
// 不要更改此文件名
// 在这里写你的事件。

package main

import (
	"github.com/dollarkillerx/govcl/vcl"
	"github.com/dollarkillerx/govcl/vcl/types"
)

//::private::
type TNewTencentTranslateEngineFields struct {
}

func (f *TNewTencentTranslateEngine) OnFormCreate(sender vcl.IObject) {

}

func (f *TNewTencentTranslateEngine) OnBtnOKClick(sender vcl.IObject) {
	if showNotAllowEmpty(f.EdtName, "名称/别名") {
		return
	}
	if showNotAllowEmpty(f.EdtSecretId, "SecretId") {
		return
	}
	if showNotAllowEmpty(f.EditSecretKey, "SecretKey") {
		return
	}

	f.SetModalResult(types.MrOk)
}