// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\electron-runtime\postprocessor\file_header\compat\file_header_1_4_2.tpl
// DO NOT EDIT!
package postprocessor

import "bytes"

func FileHeader_1_4_2(Content, Generator string, Buffer *bytes.Buffer) {
	Buffer.WriteString(`//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by Electron:`)
	Buffer.WriteString(Generator)
	Buffer.WriteString(`
//
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------
`)
	Buffer.WriteString(Content)

}
