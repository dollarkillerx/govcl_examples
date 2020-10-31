package main


import (
	// 如果你使用自定义的syso文件则不要引用此包
	_ "github.com/dollarkillerx/govcl/pkgs/winappres"
	"github.com/dollarkillerx/govcl/vcl"
)

type TMainForm struct {
	*vcl.TForm
	Btn1     *vcl.TButton
}

type TAboutForm struct {
	*vcl.TForm
	Btn1    *vcl.TButton
}

var (
	mainForm *TMainForm
	aboutForm *TAboutForm
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&mainForm)
	vcl.Application.CreateForm(&aboutForm)
	vcl.Application.Run()
}

// -- TMainForm

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("Hello")
	f.Btn1 = vcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetBounds(10, 10, 88, 28)
	f.Btn1.SetCaption("Button1")
	f.Btn1.SetOnClick(f.OnButtonClick)
}

func (f *TMainForm) OnButtonClick(sender vcl.IObject) {
	aboutForm.Show()
}


// -- TAboutForm

func (f *TAboutForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("Hello")
	f.Btn1 = vcl.NewButton(f)
	//f.Btn1.SetName("Btn1")
	f.Btn1.SetParent(f)
	f.Btn1.SetBounds(10, 10, 88, 28)
	f.Btn1.SetCaption("Button1")
	f.Btn1.SetOnClick(f.OnButtonClick)
}

func (f *TAboutForm) OnButtonClick(sender vcl.IObject) {
	vcl.ShowMessage("Hello!")
}

func (f *TAboutForm) TCloseEvent(sender vcl.IObject) {
	vcl.ShowMessage("Over")
}