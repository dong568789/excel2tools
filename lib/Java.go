package lib

import (
	"fmt"
	"github.com/dong568789/excel2tools/utils"
	"strings"
)

type JField struct {
	Name string
	Type string
	Mark string
}

type JavaClass struct {
	Name string
	Fs   []*JField
}

func (jf *JField) BuildAttribute() string {
	return fmt.Sprintf("    //%s\n	public %s %s\n", jf.parseMark(), jf.converType(), jf.Name)
}

func (jf *JField) BuildFun() string {
	types := jf.converType()
	upName := utils.ToUcFirst(jf.Name)
	getFun := fmt.Sprintf("	public %s get%s(){return this.%s;}\n", types, upName, jf.Name)
	setFun := fmt.Sprintf("	public void set%s(%s %s){this.%s = %s;}\n", upName, types, jf.Name, jf.Name, jf.Name)
	return getFun + setFun
}

func (jc *JavaClass) Build() string {
	var javaClass strings.Builder
	var attribute strings.Builder
	var fun strings.Builder
	javaClass.WriteString(fmt.Sprintf("public class %s {\n", jc.Name))
	for _, field := range jc.Fs {
		attribute.WriteString(field.BuildAttribute())
		fun.WriteString(field.BuildFun())
	}
	javaClass.WriteString(attribute.String())
	javaClass.WriteString(fun.String())

	javaClass.WriteString("}")
	return javaClass.String()
}

func (jf *JField) parseMark() string {
	return strings.Replace(jf.Mark, "\n", "|", -1)
}

func (jf *JField) converType() string {
	switch jf.Type {
	case "string":
		return "String"
	case "int":
		return "Integer"
	default:
		return "String"
	}
}
