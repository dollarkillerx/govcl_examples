// Automatically generated by the res2go, do not edit.
package main

import (
	"github.com/dollarkillerx/govcl/vcl"
)

type TNewTencentTranslateEngine struct {
	*vcl.TForm
	Label1        *vcl.TLabel
	Label2        *vcl.TLabel
	Label3        *vcl.TLabel
	Label5        *vcl.TLabel
	EdtName       *vcl.TEdit
	EdtSecretId   *vcl.TEdit
	EditSecretKey *vcl.TEdit
	BtnOK         *vcl.TButton
	Button2       *vcl.TButton

	//::private::
	TNewTencentTranslateEngineFields
}

var NewTencentTranslateEngine *TNewTencentTranslateEngine

// vcl.Application.CreateForm(&NewTencentTranslateEngine)

func NewNewTencentTranslateEngine(owner vcl.IComponent) (root *TNewTencentTranslateEngine) {
	vcl.CreateResForm(owner, &root)
	return
}

var newTencentTranslateEngineBytes = []byte("\x54\x50\x46\x30\x1A\x54\x4E\x65\x77\x54\x65\x6E\x63\x65\x6E\x74\x54\x72\x61\x6E\x73\x6C\x61\x74\x65\x45\x6E\x67\x69\x6E\x65\x19\x4E\x65\x77\x54\x65\x6E\x63\x65\x6E\x74\x54\x72\x61\x6E\x73\x6C\x61\x74\x65\x45\x6E\x67\x69\x6E\x65\x04\x4C\x65\x66\x74\x03\xE1\x03\x06\x48\x65\x69\x67\x68\x74\x03\xD9\x00\x03\x54\x6F\x70\x03\xBA\x01\x05\x57\x69\x64\x74\x68\x03\xDF\x01\x0B\x42\x6F\x72\x64\x65\x72\x49\x63\x6F\x6E\x73\x0B\x0C\x62\x69\x53\x79\x73\x74\x65\x6D\x4D\x65\x6E\x75\x00\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x08\x62\x73\x53\x69\x6E\x67\x6C\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x21\xE6\x96\xB0\xE5\xBB\xBA\xE7\xBF\xBB\xE8\xAF\x91\xE5\xBC\x95\xE6\x93\x8E\xEF\xBC\x88\xE8\x85\xBE\xE8\xAE\xAF\xE4\xBA\x91\xEF\xBC\x89\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xD9\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xDF\x01\x08\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x0E\x70\x6F\x53\x63\x72\x65\x65\x6E\x43\x65\x6E\x74\x65\x72\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x08\x32\x2E\x30\x2E\x31\x30\x2E\x30\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x16\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x18\x05\x57\x69\x64\x74\x68\x02\x41\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x10\xE5\x90\x8D\xE7\xA7\xB0\x2F\xE5\x88\xAB\xE5\x90\x8D\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x16\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x39\x05\x57\x69\x64\x74\x68\x02\x3C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0B\x53\x65\x63\x72\x65\x74\x49\x64\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x33\x04\x4C\x65\x66\x74\x02\x16\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x5B\x05\x57\x69\x64\x74\x68\x02\x45\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\x53\x65\x63\x72\x65\x74\x4B\x65\x79\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x35\x04\x4C\x65\x66\x74\x02\x17\x06\x48\x65\x69\x67\x68\x74\x02\x22\x03\x54\x6F\x70\x02\x7E\x05\x57\x69\x64\x74\x68\x03\xBE\x00\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x3E\xE8\xAF\xB4\xE6\x98\x8E\xEF\xBC\x9A\x0D\x0A\xE8\xAF\xB7\xE5\xA1\xAB\xE5\x86\x99\xE2\x80\x9C\xE8\x85\xBE\xE8\xAE\xAF\xE4\xBA\x91\xE2\x80\x9D\xE5\x88\x9B\xE5\xBB\xBA\xE7\x9A\x84\xE5\xAD\x90\xE7\x94\xA8\xE6\x88\xB7\xE5\xAF\x86\xE9\x92\xA5\xE3\x80\x82\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x07\x45\x64\x74\x4E\x61\x6D\x65\x04\x4C\x65\x66\x74\x02\x6F\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x15\x05\x57\x69\x64\x74\x68\x03\x61\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x05\x54\x45\x64\x69\x74\x0B\x45\x64\x74\x53\x65\x63\x72\x65\x74\x49\x64\x04\x4C\x65\x66\x74\x02\x6F\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x34\x05\x57\x69\x64\x74\x68\x03\x61\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x05\x54\x45\x64\x69\x74\x0D\x45\x64\x69\x74\x53\x65\x63\x72\x65\x74\x4B\x65\x79\x04\x4C\x65\x66\x74\x02\x6F\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x56\x05\x57\x69\x64\x74\x68\x03\x61\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x05\x42\x74\x6E\x4F\x4B\x04\x4C\x65\x66\x74\x03\x2F\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xB0\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE7\xA1\xAE\xE5\xAE\x9A\xE6\x96\xB0\xE5\xA2\x9E\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x32\x04\x4C\x65\x66\x74\x03\x82\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xB0\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x8F\x96\xE6\xB6\x88\x0B\x4D\x6F\x64\x61\x6C\x52\x65\x73\x75\x6C\x74\x02\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x00\x00")

// Register Form Resources
var _ = vcl.RegisterFormResource(NewTencentTranslateEngine, &newTencentTranslateEngineBytes)
