// Automatically generated by the res2go, do not edit.

package main

import (
	"github.com/dollarkillerx/govcl/vcl"
)

type TGdipForm struct {
	*vcl.TForm

	//::private::
	TGdipFormFields
}

var GdipForm *TGdipForm

// Loaded in bytes.
// vcl.Application.CreateForm(&GdipForm)

func NewGdipForm(owner vcl.IComponent) (root *TGdipForm) {
	vcl.CreateResForm(owner, &root)
	return
}

var GdipFormBytes = []byte("\x54\x50\x46\x30\x0B\x54\x44\x65\x73\x69\x67\x6E\x46\x6F\x72\x6D\x08\x47\x64\x69\x70\x46\x6F\x72\x6D\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x03\x3E\x02\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\x6A\x03\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x06\x62\x73\x4E\x6F\x6E\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0A\x47\x44\x49\x2B\xE6\xB5\x8B\xE8\xAF\x95\x05\x43\x6F\x6C\x6F\x72\x07\x09\x63\x6C\x42\x74\x6E\x46\x61\x63\x65\x0E\x44\x6F\x75\x62\x6C\x65\x42\x75\x66\x66\x65\x72\x65\x64\x09\x0A\x46\x6F\x6E\x74\x2E\x43\x6F\x6C\x6F\x72\x07\x0C\x63\x6C\x57\x69\x6E\x64\x6F\x77\x54\x65\x78\x74\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xF3\x09\x46\x6F\x6E\x74\x2E\x4E\x61\x6D\x65\x06\x06\x54\x61\x68\x6F\x6D\x61\x14\x50\x61\x72\x65\x6E\x74\x44\x6F\x75\x62\x6C\x65\x42\x75\x66\x66\x65\x72\x65\x64\x08\x08\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x0E\x70\x6F\x53\x63\x72\x65\x65\x6E\x43\x65\x6E\x74\x65\x72\x00\x00")

// Register Form Resources
var _ = vcl.RegisterFormResource(GdipForm, &GdipFormBytes)
