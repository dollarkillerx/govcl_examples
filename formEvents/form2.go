// Automatically generated by the res2go, do not edit.

package main

import (
	"github.com/dollarkillerx/govcl/vcl"
)

type TForm2 struct {
	*vcl.TForm
	RadioGroup1 *vcl.TRadioGroup
	Button1     *vcl.TButton
	Button2     *vcl.TButton

	//::private::
	TForm2Fields
}

var Form2 *TForm2

// Loaded in bytes.
// vcl.Application.CreateForm(&Form2)

func NewForm2(owner vcl.IComponent) (root *TForm2) {
	vcl.CreateResForm(owner, &root)
	return
}

var Form2Bytes = []byte("\x54\x50\x46\x30\x0B\x54\x44\x65\x73\x69\x67\x6E\x46\x6F\x72\x6D\x05\x46\x6F\x72\x6D\x32\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x03\xC1\x00\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\x2E\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x05\x46\x6F\x72\x6D\x32\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xC1\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x2E\x01\x05\x43\x6F\x6C\x6F\x72\x07\x09\x63\x6C\x42\x74\x6E\x46\x61\x63\x65\x0A\x46\x6F\x6E\x74\x2E\x43\x6F\x6C\x6F\x72\x07\x0C\x63\x6C\x57\x69\x6E\x64\x6F\x77\x54\x65\x78\x74\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xF3\x09\x46\x6F\x6E\x74\x2E\x4E\x61\x6D\x65\x06\x06\x54\x61\x68\x6F\x6D\x61\x08\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x10\x70\x6F\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x43\x65\x6E\x74\x65\x72\x00\x0B\x54\x52\x61\x64\x69\x6F\x47\x72\x6F\x75\x70\x0B\x52\x61\x64\x69\x6F\x47\x72\x6F\x75\x70\x31\x04\x4C\x65\x66\x74\x02\x12\x06\x48\x65\x69\x67\x68\x74\x02\x7E\x03\x54\x6F\x70\x02\x06\x05\x57\x69\x64\x74\x68\x03\x0C\x01\x08\x41\x75\x74\x6F\x46\x69\x6C\x6C\x09\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0B\x52\x61\x64\x69\x6F\x47\x72\x6F\x75\x70\x31\x1C\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x4C\x65\x66\x74\x52\x69\x67\x68\x74\x53\x70\x61\x63\x69\x6E\x67\x02\x06\x1D\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x45\x6E\x6C\x61\x72\x67\x65\x48\x6F\x72\x69\x7A\x6F\x6E\x74\x61\x6C\x07\x18\x63\x72\x73\x48\x6F\x6D\x6F\x67\x65\x6E\x6F\x75\x73\x43\x68\x69\x6C\x64\x52\x65\x73\x69\x7A\x65\x1B\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x45\x6E\x6C\x61\x72\x67\x65\x56\x65\x72\x74\x69\x63\x61\x6C\x07\x18\x63\x72\x73\x48\x6F\x6D\x6F\x67\x65\x6E\x6F\x75\x73\x43\x68\x69\x6C\x64\x52\x65\x73\x69\x7A\x65\x1C\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x53\x68\x72\x69\x6E\x6B\x48\x6F\x72\x69\x7A\x6F\x6E\x74\x61\x6C\x07\x0E\x63\x72\x73\x53\x63\x61\x6C\x65\x43\x68\x69\x6C\x64\x73\x1A\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x53\x68\x72\x69\x6E\x6B\x56\x65\x72\x74\x69\x63\x61\x6C\x07\x0E\x63\x72\x73\x53\x63\x61\x6C\x65\x43\x68\x69\x6C\x64\x73\x12\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x4C\x61\x79\x6F\x75\x74\x07\x1D\x63\x63\x6C\x4C\x65\x66\x74\x54\x6F\x52\x69\x67\x68\x74\x54\x68\x65\x6E\x54\x6F\x70\x54\x6F\x42\x6F\x74\x74\x6F\x6D\x1B\x43\x68\x69\x6C\x64\x53\x69\x7A\x69\x6E\x67\x2E\x43\x6F\x6E\x74\x72\x6F\x6C\x73\x50\x65\x72\x4C\x69\x6E\x65\x02\x01\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x69\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x08\x01\x0D\x49\x74\x65\x6D\x73\x2E\x53\x74\x72\x69\x6E\x67\x73\x01\x06\x01\x31\x06\x01\x32\x06\x01\x33\x06\x01\x34\x06\x01\x35\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x31\x04\x4C\x65\x66\x74\x02\x72\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x92\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x02\x4F\x4B\x0B\x4D\x6F\x64\x61\x6C\x52\x65\x73\x75\x6C\x74\x02\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x32\x04\x4C\x65\x66\x74\x03\xC3\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\x92\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\x43\x61\x6E\x63\x65\x6C\x0B\x4D\x6F\x64\x61\x6C\x52\x65\x73\x75\x6C\x74\x02\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x00")

// 注册窗口资源
var _ = vcl.RegisterFormResource(Form2, &Form2Bytes)
