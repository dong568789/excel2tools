package lib

import (
	"fmt"
	"github.com/dong568789/excel2tools/utils"
	"strings"
)

type GField struct {
	Name string
	Type string
	Mark string
}

type Local struct {
	Name string
	Fs   []*GField
}

func (gf *GField) BuildAttribute() string {
	return fmt.Sprintf("	%s %s  //%s\n", utils.ToUcFirst(gf.Name), gf.converType(), gf.parseMark())
}

func (l *Local) Build() string {
	var attribute strings.Builder
	var protobufStr strings.Builder
	protobufStr.WriteString(fmt.Sprintf("package models\n\n"))
	protobufStr.WriteString(fmt.Sprintf("type %s struct {\n", utils.ToUcFirst(l.Name)))
	for _, m := range l.Fs {
		attribute.WriteString(m.BuildAttribute())
	}
	protobufStr.WriteString(attribute.String())
	protobufStr.WriteString(fmt.Sprintf("}\n"))
	return protobufStr.String()
}

func (gf *GField) parseMark() string {
	return strings.Replace(gf.Mark, "\n", "|", -1)
}

func (gf *GField) converType() string {
	switch gf.Type {
	case "string":
		return "string"
	case "int":
		return "int"
	default:
		return "string"
	}
}
