// 由GOVCL UI设计器自动生成，不要编辑
package main

import (
	"github.com/dollarkillerx/govcl/vcl"
)

type TMainForm struct {
	*vcl.TForm
	Panel1            *vcl.TPanel
	StatusBar1        *vcl.TStatusBar
	Panel2            *vcl.TPanel
	PnlWeb            *vcl.TPanel
	Panel3            *vcl.TPanel
	BtnBack           *vcl.TButton
	BtnForward        *vcl.TButton
	BtnReload         *vcl.TButton
	Panel4            *vcl.TPanel
	BtnNav            *vcl.TButton
	Panel5            *vcl.TPanel
	Label1            *vcl.TLabel
	EdtURL            *vcl.TEdit
	BtnLoadFromFile   *vcl.TButton
	BtnLoadFromString *vcl.TButton
	OpenDialog1       *vcl.TOpenDialog
	ActionList1       *vcl.TActionList
	ActGoBack         *vcl.TAction
	ActGoForward      *vcl.TAction
	ActNav            *vcl.TAction
	BtnExecJS         *vcl.TButton

	//::private::
	TMainFormFields
}

var MainForm *TMainForm

// 以字节形式加载
// vcl.Application.CreateForm(&MainForm)

var (
	mainFormBytes = []byte{
		0x54, 0x50, 0x46, 0x30, 0x05, 0x54, 0x46, 0x6F, 0x72, 0x6D, 0x08, 0x4D,
		0x61, 0x69, 0x6E, 0x46, 0x6F, 0x72, 0x6D, 0x04, 0x4C, 0x65, 0x66, 0x74,
		0x02, 0x00, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x00, 0x07, 0x43, 0x61, 0x70,
		0x74, 0x69, 0x6F, 0x6E, 0x14, 0x0F, 0x00, 0x00, 0x00, 0x4D, 0x69, 0x6E,
		0x69, 0x62, 0x6C, 0x69, 0x6E, 0x6B, 0xE6, 0xB5, 0x8B, 0xE8, 0xAF, 0x95,
		0x0C, 0x43, 0x6C, 0x69, 0x65, 0x6E, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68,
		0x74, 0x03, 0x9D, 0x02, 0x0B, 0x43, 0x6C, 0x69, 0x65, 0x6E, 0x74, 0x57,
		0x69, 0x64, 0x74, 0x68, 0x03, 0x90, 0x04, 0x05, 0x43, 0x6F, 0x6C, 0x6F,
		0x72, 0x07, 0x09, 0x63, 0x6C, 0x42, 0x74, 0x6E, 0x46, 0x61, 0x63, 0x65,
		0x0C, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x43, 0x68, 0x61, 0x72, 0x73, 0x65,
		0x74, 0x07, 0x0F, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4C, 0x54, 0x5F, 0x43,
		0x48, 0x41, 0x52, 0x53, 0x45, 0x54, 0x0A, 0x46, 0x6F, 0x6E, 0x74, 0x2E,
		0x43, 0x6F, 0x6C, 0x6F, 0x72, 0x07, 0x0C, 0x63, 0x6C, 0x57, 0x69, 0x6E,
		0x64, 0x6F, 0x77, 0x54, 0x65, 0x78, 0x74, 0x0B, 0x46, 0x6F, 0x6E, 0x74,
		0x2E, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0xF3, 0x09, 0x46, 0x6F,
		0x6E, 0x74, 0x2E, 0x4E, 0x61, 0x6D, 0x65, 0x06, 0x06, 0x54, 0x61, 0x68,
		0x6F, 0x6D, 0x61, 0x0A, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x53, 0x74, 0x79,
		0x6C, 0x65, 0x0B, 0x00, 0x0A, 0x4B, 0x65, 0x79, 0x50, 0x72, 0x65, 0x76,
		0x69, 0x65, 0x77, 0x09, 0x0E, 0x4F, 0x6C, 0x64, 0x43, 0x72, 0x65, 0x61,
		0x74, 0x65, 0x4F, 0x72, 0x64, 0x65, 0x72, 0x08, 0x08, 0x50, 0x6F, 0x73,
		0x69, 0x74, 0x69, 0x6F, 0x6E, 0x07, 0x0E, 0x70, 0x6F, 0x53, 0x63, 0x72,
		0x65, 0x65, 0x6E, 0x43, 0x65, 0x6E, 0x74, 0x65, 0x72, 0x0D, 0x50, 0x69,
		0x78, 0x65, 0x6C, 0x73, 0x50, 0x65, 0x72, 0x49, 0x6E, 0x63, 0x68, 0x02,
		0x60, 0x0A, 0x54, 0x65, 0x78, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
		0x02, 0x10, 0x00, 0x06, 0x54, 0x50, 0x61, 0x6E, 0x65, 0x6C, 0x06, 0x50,
		0x61, 0x6E, 0x65, 0x6C, 0x31, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x00,
		0x03, 0x54, 0x6F, 0x70, 0x02, 0x00, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68,
		0x03, 0x90, 0x04, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x35,
		0x05, 0x41, 0x6C, 0x69, 0x67, 0x6E, 0x07, 0x05, 0x61, 0x6C, 0x54, 0x6F,
		0x70, 0x0A, 0x42, 0x65, 0x76, 0x65, 0x6C, 0x4F, 0x75, 0x74, 0x65, 0x72,
		0x07, 0x06, 0x62, 0x76, 0x4E, 0x6F, 0x6E, 0x65, 0x08, 0x54, 0x61, 0x62,
		0x4F, 0x72, 0x64, 0x65, 0x72, 0x02, 0x00, 0x0C, 0x45, 0x78, 0x70, 0x6C,
		0x69, 0x63, 0x69, 0x74, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x01, 0x00, 0x06,
		0x54, 0x50, 0x61, 0x6E, 0x65, 0x6C, 0x06, 0x50, 0x61, 0x6E, 0x65, 0x6C,
		0x33, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x00, 0x03, 0x54, 0x6F, 0x70,
		0x02, 0x00, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x03, 0xB9, 0x00, 0x06,
		0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x35, 0x05, 0x41, 0x6C, 0x69,
		0x67, 0x6E, 0x07, 0x06, 0x61, 0x6C, 0x4C, 0x65, 0x66, 0x74, 0x0A, 0x42,
		0x65, 0x76, 0x65, 0x6C, 0x4F, 0x75, 0x74, 0x65, 0x72, 0x07, 0x06, 0x62,
		0x76, 0x4E, 0x6F, 0x6E, 0x65, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64,
		0x65, 0x72, 0x02, 0x00, 0x0B, 0x45, 0x78, 0x70, 0x6C, 0x69, 0x63, 0x69,
		0x74, 0x54, 0x6F, 0x70, 0x02, 0x01, 0x0E, 0x45, 0x78, 0x70, 0x6C, 0x69,
		0x63, 0x69, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x33, 0x00,
		0x07, 0x54, 0x42, 0x75, 0x74, 0x74, 0x6F, 0x6E, 0x07, 0x42, 0x74, 0x6E,
		0x42, 0x61, 0x63, 0x6B, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x1C, 0x03,
		0x54, 0x6F, 0x70, 0x02, 0x0F, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x02,
		0x23, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x19, 0x06, 0x41,
		0x63, 0x74, 0x69, 0x6F, 0x6E, 0x07, 0x09, 0x41, 0x63, 0x74, 0x47, 0x6F,
		0x42, 0x61, 0x63, 0x6B, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65,
		0x72, 0x02, 0x00, 0x00, 0x00, 0x07, 0x54, 0x42, 0x75, 0x74, 0x74, 0x6F,
		0x6E, 0x0A, 0x42, 0x74, 0x6E, 0x46, 0x6F, 0x72, 0x77, 0x61, 0x72, 0x64,
		0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x49, 0x03, 0x54, 0x6F, 0x70, 0x02,
		0x0F, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x02, 0x23, 0x06, 0x48, 0x65,
		0x69, 0x67, 0x68, 0x74, 0x02, 0x19, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6F,
		0x6E, 0x07, 0x0C, 0x41, 0x63, 0x74, 0x47, 0x6F, 0x46, 0x6F, 0x72, 0x77,
		0x61, 0x72, 0x64, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65, 0x72,
		0x02, 0x01, 0x00, 0x00, 0x07, 0x54, 0x42, 0x75, 0x74, 0x74, 0x6F, 0x6E,
		0x09, 0x42, 0x74, 0x6E, 0x52, 0x65, 0x6C, 0x6F, 0x61, 0x64, 0x04, 0x4C,
		0x65, 0x66, 0x74, 0x02, 0x75, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x0F, 0x05,
		0x57, 0x69, 0x64, 0x74, 0x68, 0x02, 0x23, 0x06, 0x48, 0x65, 0x69, 0x67,
		0x68, 0x74, 0x02, 0x19, 0x07, 0x43, 0x61, 0x70, 0x74, 0x69, 0x6F, 0x6E,
		0x06, 0x01, 0x4F, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65, 0x72,
		0x02, 0x02, 0x00, 0x00, 0x00, 0x06, 0x54, 0x50, 0x61, 0x6E, 0x65, 0x6C,
		0x06, 0x50, 0x61, 0x6E, 0x65, 0x6C, 0x34, 0x04, 0x4C, 0x65, 0x66, 0x74,
		0x03, 0x14, 0x04, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x00, 0x05, 0x57, 0x69,
		0x64, 0x74, 0x68, 0x02, 0x7C, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
		0x02, 0x35, 0x05, 0x41, 0x6C, 0x69, 0x67, 0x6E, 0x07, 0x07, 0x61, 0x6C,
		0x52, 0x69, 0x67, 0x68, 0x74, 0x0A, 0x42, 0x65, 0x76, 0x65, 0x6C, 0x4F,
		0x75, 0x74, 0x65, 0x72, 0x07, 0x06, 0x62, 0x76, 0x4E, 0x6F, 0x6E, 0x65,
		0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65, 0x72, 0x02, 0x01, 0x0C,
		0x45, 0x78, 0x70, 0x6C, 0x69, 0x63, 0x69, 0x74, 0x4C, 0x65, 0x66, 0x74,
		0x03, 0x12, 0x04, 0x0B, 0x45, 0x78, 0x70, 0x6C, 0x69, 0x63, 0x69, 0x74,
		0x54, 0x6F, 0x70, 0x02, 0x01, 0x0E, 0x45, 0x78, 0x70, 0x6C, 0x69, 0x63,
		0x69, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x33, 0x00, 0x07,
		0x54, 0x42, 0x75, 0x74, 0x74, 0x6F, 0x6E, 0x06, 0x42, 0x74, 0x6E, 0x4E,
		0x61, 0x76, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x0E, 0x03, 0x54, 0x6F,
		0x70, 0x02, 0x0D, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x02, 0x4B, 0x06,
		0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x19, 0x06, 0x41, 0x63, 0x74,
		0x69, 0x6F, 0x6E, 0x07, 0x06, 0x41, 0x63, 0x74, 0x4E, 0x61, 0x76, 0x08,
		0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65, 0x72, 0x02, 0x00, 0x00, 0x00,
		0x00, 0x06, 0x54, 0x50, 0x61, 0x6E, 0x65, 0x6C, 0x06, 0x50, 0x61, 0x6E,
		0x65, 0x6C, 0x35, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x03, 0xB9, 0x00, 0x03,
		0x54, 0x6F, 0x70, 0x02, 0x00, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x03,
		0x5B, 0x03, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x35, 0x05,
		0x41, 0x6C, 0x69, 0x67, 0x6E, 0x07, 0x08, 0x61, 0x6C, 0x43, 0x6C, 0x69,
		0x65, 0x6E, 0x74, 0x0A, 0x42, 0x65, 0x76, 0x65, 0x6C, 0x4F, 0x75, 0x74,
		0x65, 0x72, 0x07, 0x06, 0x62, 0x76, 0x4E, 0x6F, 0x6E, 0x65, 0x08, 0x54,
		0x61, 0x62, 0x4F, 0x72, 0x64, 0x65, 0x72, 0x02, 0x02, 0x0C, 0x45, 0x78,
		0x70, 0x6C, 0x69, 0x63, 0x69, 0x74, 0x4C, 0x65, 0x66, 0x74, 0x03, 0x9B,
		0x01, 0x0B, 0x45, 0x78, 0x70, 0x6C, 0x69, 0x63, 0x69, 0x74, 0x54, 0x6F,
		0x70, 0x02, 0x0A, 0x0D, 0x45, 0x78, 0x70, 0x6C, 0x69, 0x63, 0x69, 0x74,
		0x57, 0x69, 0x64, 0x74, 0x68, 0x03, 0xB9, 0x00, 0x0E, 0x45, 0x78, 0x70,
		0x6C, 0x69, 0x63, 0x69, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02,
		0x29, 0x0A, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6E, 0x53, 0x69, 0x7A, 0x65,
		0x01, 0x03, 0x5B, 0x03, 0x02, 0x35, 0x00, 0x00, 0x06, 0x54, 0x4C, 0x61,
		0x62, 0x65, 0x6C, 0x06, 0x4C, 0x61, 0x62, 0x65, 0x6C, 0x31, 0x04, 0x4C,
		0x65, 0x66, 0x74, 0x02, 0x00, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x00, 0x05,
		0x57, 0x69, 0x64, 0x74, 0x68, 0x02, 0x30, 0x06, 0x48, 0x65, 0x69, 0x67,
		0x68, 0x74, 0x02, 0x35, 0x05, 0x41, 0x6C, 0x69, 0x67, 0x6E, 0x07, 0x06,
		0x61, 0x6C, 0x4C, 0x65, 0x66, 0x74, 0x07, 0x43, 0x61, 0x70, 0x74, 0x69,
		0x6F, 0x6E, 0x12, 0x03, 0x00, 0x00, 0x00, 0x30, 0x57, 0x40, 0x57, 0x1A,
		0xFF, 0x06, 0x4C, 0x61, 0x79, 0x6F, 0x75, 0x74, 0x07, 0x08, 0x74, 0x6C,
		0x43, 0x65, 0x6E, 0x74, 0x65, 0x72, 0x0B, 0x45, 0x78, 0x70, 0x6C, 0x69,
		0x63, 0x69, 0x74, 0x54, 0x6F, 0x70, 0x02, 0xFF, 0x0E, 0x45, 0x78, 0x70,
		0x6C, 0x69, 0x63, 0x69, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02,
		0x33, 0x00, 0x00, 0x05, 0x54, 0x45, 0x64, 0x69, 0x74, 0x06, 0x45, 0x64,
		0x74, 0x55, 0x52, 0x4C, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x37, 0x03,
		0x54, 0x6F, 0x70, 0x02, 0x0E, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x03,
		0x11, 0x03, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x18, 0x07,
		0x41, 0x6E, 0x63, 0x68, 0x6F, 0x72, 0x73, 0x0B, 0x06, 0x61, 0x6B, 0x4C,
		0x65, 0x66, 0x74, 0x05, 0x61, 0x6B, 0x54, 0x6F, 0x70, 0x07, 0x61, 0x6B,
		0x52, 0x69, 0x67, 0x68, 0x74, 0x00, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72,
		0x64, 0x65, 0x72, 0x02, 0x00, 0x04, 0x54, 0x65, 0x78, 0x74, 0x06, 0x1F,
		0x68, 0x74, 0x74, 0x70, 0x73, 0x3A, 0x2F, 0x2F, 0x67, 0x69, 0x74, 0x68,
		0x75, 0x62, 0x2E, 0x63, 0x6F, 0x6D, 0x2F, 0x79, 0x69, 0x6E, 0x67, 0x33,
		0x32, 0x2F, 0x67, 0x6F, 0x76, 0x63, 0x6C, 0x08, 0x54, 0x65, 0x78, 0x74,
		0x48, 0x69, 0x6E, 0x74, 0x14, 0x25, 0x00, 0x00, 0x00, 0xE4, 0xBE, 0x8B,
		0xEF, 0xBC, 0x9A, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3A, 0x2F, 0x2F, 0x67,
		0x69, 0x74, 0x68, 0x75, 0x62, 0x2E, 0x63, 0x6F, 0x6D, 0x2F, 0x79, 0x69,
		0x6E, 0x67, 0x33, 0x32, 0x2F, 0x67, 0x6F, 0x76, 0x63, 0x6C, 0x0D, 0x45,
		0x78, 0x70, 0x6C, 0x69, 0x63, 0x69, 0x74, 0x57, 0x69, 0x64, 0x74, 0x68,
		0x03, 0x0F, 0x03, 0x00, 0x00, 0x00, 0x00, 0x0A, 0x54, 0x53, 0x74, 0x61,
		0x74, 0x75, 0x73, 0x42, 0x61, 0x72, 0x0A, 0x53, 0x74, 0x61, 0x74, 0x75,
		0x73, 0x42, 0x61, 0x72, 0x31, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x00,
		0x03, 0x54, 0x6F, 0x70, 0x03, 0x8A, 0x02, 0x05, 0x57, 0x69, 0x64, 0x74,
		0x68, 0x03, 0x90, 0x04, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02,
		0x13, 0x06, 0x50, 0x61, 0x6E, 0x65, 0x6C, 0x73, 0x0E, 0x00, 0x0B, 0x45,
		0x78, 0x70, 0x6C, 0x69, 0x63, 0x69, 0x74, 0x54, 0x6F, 0x70, 0x03, 0xA2,
		0x01, 0x00, 0x00, 0x06, 0x54, 0x50, 0x61, 0x6E, 0x65, 0x6C, 0x06, 0x50,
		0x61, 0x6E, 0x65, 0x6C, 0x32, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x00,
		0x03, 0x54, 0x6F, 0x70, 0x02, 0x35, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68,
		0x03, 0x90, 0x04, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x29,
		0x05, 0x41, 0x6C, 0x69, 0x67, 0x6E, 0x07, 0x05, 0x61, 0x6C, 0x54, 0x6F,
		0x70, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65, 0x72, 0x02, 0x02,
		0x0B, 0x45, 0x78, 0x70, 0x6C, 0x69, 0x63, 0x69, 0x74, 0x54, 0x6F, 0x70,
		0x02, 0x34, 0x00, 0x07, 0x54, 0x42, 0x75, 0x74, 0x74, 0x6F, 0x6E, 0x0F,
		0x42, 0x74, 0x6E, 0x4C, 0x6F, 0x61, 0x64, 0x46, 0x72, 0x6F, 0x6D, 0x46,
		0x69, 0x6C, 0x65, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x34, 0x03, 0x54,
		0x6F, 0x70, 0x02, 0x08, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x02, 0x7E,
		0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x19, 0x07, 0x43, 0x61,
		0x70, 0x74, 0x69, 0x6F, 0x6E, 0x12, 0x07, 0x00, 0x00, 0x00, 0x4B, 0x6D,
		0xD5, 0x8B, 0xCE, 0x4E, 0x87, 0x65, 0xF6, 0x4E, 0xA0, 0x52, 0x7D, 0x8F,
		0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65, 0x72, 0x02, 0x00, 0x00,
		0x00, 0x07, 0x54, 0x42, 0x75, 0x74, 0x74, 0x6F, 0x6E, 0x11, 0x42, 0x74,
		0x6E, 0x4C, 0x6F, 0x61, 0x64, 0x46, 0x72, 0x6F, 0x6D, 0x53, 0x74, 0x72,
		0x69, 0x6E, 0x67, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x03, 0xBE, 0x00, 0x03,
		0x54, 0x6F, 0x70, 0x02, 0x08, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x03,
		0x90, 0x00, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x19, 0x07,
		0x43, 0x61, 0x70, 0x74, 0x69, 0x6F, 0x6E, 0x12, 0x08, 0x00, 0x00, 0x00,
		0x4B, 0x6D, 0xD5, 0x8B, 0xCE, 0x4E, 0x57, 0x5B, 0x26, 0x7B, 0x32, 0x4E,
		0xA0, 0x52, 0x7D, 0x8F, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65,
		0x72, 0x02, 0x01, 0x00, 0x00, 0x07, 0x54, 0x42, 0x75, 0x74, 0x74, 0x6F,
		0x6E, 0x09, 0x42, 0x74, 0x6E, 0x45, 0x78, 0x65, 0x63, 0x4A, 0x53, 0x04,
		0x4C, 0x65, 0x66, 0x74, 0x03, 0x60, 0x01, 0x03, 0x54, 0x6F, 0x70, 0x02,
		0x08, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x02, 0x4B, 0x06, 0x48, 0x65,
		0x69, 0x67, 0x68, 0x74, 0x02, 0x19, 0x07, 0x43, 0x61, 0x70, 0x74, 0x69,
		0x6F, 0x6E, 0x12, 0x04, 0x00, 0x00, 0x00, 0x67, 0x62, 0x4C, 0x88, 0x4A,
		0x00, 0x73, 0x00, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65, 0x72,
		0x02, 0x02, 0x00, 0x00, 0x00, 0x06, 0x54, 0x50, 0x61, 0x6E, 0x65, 0x6C,
		0x06, 0x50, 0x6E, 0x6C, 0x57, 0x65, 0x62, 0x04, 0x4C, 0x65, 0x66, 0x74,
		0x02, 0x00, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x5E, 0x05, 0x57, 0x69, 0x64,
		0x74, 0x68, 0x03, 0x90, 0x04, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
		0x03, 0x2C, 0x02, 0x05, 0x41, 0x6C, 0x69, 0x67, 0x6E, 0x07, 0x08, 0x61,
		0x6C, 0x43, 0x6C, 0x69, 0x65, 0x6E, 0x74, 0x0A, 0x42, 0x65, 0x76, 0x65,
		0x6C, 0x4F, 0x75, 0x74, 0x65, 0x72, 0x07, 0x06, 0x62, 0x76, 0x4E, 0x6F,
		0x6E, 0x65, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65, 0x72, 0x02,
		0x03, 0x0C, 0x45, 0x78, 0x70, 0x6C, 0x69, 0x63, 0x69, 0x74, 0x4C, 0x65,
		0x66, 0x74, 0x03, 0x8D, 0x00, 0x0B, 0x45, 0x78, 0x70, 0x6C, 0x69, 0x63,
		0x69, 0x74, 0x54, 0x6F, 0x70, 0x03, 0xB1, 0x00, 0x0D, 0x45, 0x78, 0x70,
		0x6C, 0x69, 0x63, 0x69, 0x74, 0x57, 0x69, 0x64, 0x74, 0x68, 0x03, 0xB9,
		0x00, 0x0E, 0x45, 0x78, 0x70, 0x6C, 0x69, 0x63, 0x69, 0x74, 0x48, 0x65,
		0x69, 0x67, 0x68, 0x74, 0x02, 0x29, 0x00, 0x00, 0x0B, 0x54, 0x4F, 0x70,
		0x65, 0x6E, 0x44, 0x69, 0x61, 0x6C, 0x6F, 0x67, 0x0B, 0x4F, 0x70, 0x65,
		0x6E, 0x44, 0x69, 0x61, 0x6C, 0x6F, 0x67, 0x31, 0x04, 0x4C, 0x65, 0x66,
		0x74, 0x03, 0x90, 0x00, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x1D, 0x00, 0x00,
		0x0B, 0x54, 0x41, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x4C, 0x69, 0x73, 0x74,
		0x0B, 0x41, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x4C, 0x69, 0x73, 0x74, 0x31,
		0x04, 0x4C, 0x65, 0x66, 0x74, 0x03, 0x44, 0x01, 0x03, 0x54, 0x6F, 0x70,
		0x02, 0x4D, 0x00, 0x07, 0x54, 0x41, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x09,
		0x41, 0x63, 0x74, 0x47, 0x6F, 0x42, 0x61, 0x63, 0x6B, 0x07, 0x43, 0x61,
		0x70, 0x74, 0x69, 0x6F, 0x6E, 0x06, 0x02, 0x3C, 0x2D, 0x04, 0x48, 0x69,
		0x6E, 0x74, 0x12, 0x02, 0x00, 0x00, 0x00, 0x0E, 0x54, 0x00, 0x90, 0x00,
		0x00, 0x07, 0x54, 0x41, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x0C, 0x41, 0x63,
		0x74, 0x47, 0x6F, 0x46, 0x6F, 0x72, 0x77, 0x61, 0x72, 0x64, 0x07, 0x43,
		0x61, 0x70, 0x74, 0x69, 0x6F, 0x6E, 0x06, 0x02, 0x2D, 0x3E, 0x04, 0x48,
		0x69, 0x6E, 0x74, 0x12, 0x02, 0x00, 0x00, 0x00, 0x4D, 0x52, 0xDB, 0x8F,
		0x00, 0x00, 0x07, 0x54, 0x41, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x06, 0x41,
		0x63, 0x74, 0x4E, 0x61, 0x76, 0x07, 0x43, 0x61, 0x70, 0x74, 0x69, 0x6F,
		0x6E, 0x12, 0x02, 0x00, 0x00, 0x00, 0x6C, 0x8F, 0x30, 0x52, 0x04, 0x48,
		0x69, 0x6E, 0x74, 0x12, 0x08, 0x00, 0x00, 0x00, 0xF3, 0x8D, 0x6C, 0x8F,
		0x30, 0x52, 0x07, 0x63, 0x9A, 0x5B, 0x55, 0x00, 0x52, 0x00, 0x4C, 0x00,
		0x00, 0x00, 0x00, 0x00}
)

var _ = vcl.RegisterFormResource(MainForm, &mainFormBytes)
