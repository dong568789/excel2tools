package lib

import (
	"fmt"
	"strconv"
	"strings"
)

type PField struct {
	Name string
	Type string
	Mark string
}

type Message struct {
	Name string
	Fs   []*PField
}

func (pf *PField) BuildAttribute() string {
	return fmt.Sprintf("	%s %s = [index]; //%s\n", pf.converType(), pf.Name, pf.parseMark())
}

func (m *Message) Build() string {
	var attribute strings.Builder
	var protobufStr strings.Builder
	protobufStr.WriteString(fmt.Sprintf("syntax = \"proto3\";\n"))
	protobufStr.WriteString(fmt.Sprintf("package Protobuf;\n\n"))
	protobufStr.WriteString(fmt.Sprintf("message %s {\n", m.Name))
	for k, m := range m.Fs {
		attribute.WriteString(strings.Replace(m.BuildAttribute(), "[index]", strconv.Itoa(k+1), -1))
	}
	protobufStr.WriteString(attribute.String())
	protobufStr.WriteString(fmt.Sprintf("}\n"))
	return protobufStr.String()
}

func (pf *PField) parseMark() string {
	return strings.Replace(pf.Mark, "\n", "|", -1)
}

func (pf *PField) converType() string {
	switch pf.Type {
	case "string":
		return "string"
	case "int":
		return "int32"
	default:
		return "string"
	}
}
