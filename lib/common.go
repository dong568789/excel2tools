package lib

import (
	"encoding/json"
	"fmt"
	"github.com/dong568789/excel2tools/utils"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
)

type Excel struct {
	rows      [][]string
	FieldType []string
	FieldName []string
	FieldMark []string
	FieldRead []string
	SheetName string
	output    string
}

const (
	FieldMark = iota
	FieldRead
	FieldType
	FieldName
)

func NewExcel(output string) *Excel {
	e := &Excel{
		output: output,
	}
	return e
}

func (e *Excel) Read(input string) {
	f, err := excelize.OpenFile(input)
	if err != nil {
		panic(err)
	}
	defer func() {
		f.Close()
	}()
	i := f.GetActiveSheetIndex()

	sheetName := f.GetSheetName(i)
	e.setSheetName(sheetName)

	rows, err := f.GetRows(sheetName)
	if err != nil {
		panic(err)
	}

	e.parseRows(rows)
}

func (e *Excel) Export() {
	e.ExportJson()
	e.ExportJava()
	e.ExportLocal()
	e.ExportProtobuf()
}

func (e *Excel) ExportJson() {
	path := fmt.Sprintf("%s/systemJson.json", e.output)
	utils.Write(path, e.ToJson())
}

func (e *Excel) ExportJava() {
	path := fmt.Sprintf("%s/%s.java", e.output, e.SheetName)
	utils.Write(path, e.ToJava())
}

func (e *Excel) ExportProtobuf() {
	path := fmt.Sprintf("%s/%s.proto", e.output, e.SheetName)
	utils.Write(path, e.ToProtobuf())
}

func (e *Excel) ExportLocal() {
	path := fmt.Sprintf("%s/%s.go", e.output, e.SheetName)
	utils.Write(path, e.ToLocal())
}

func (e *Excel) ToJava() []byte {
	javaClass := &JavaClass{
		Name: e.SheetName,
	}
	for k, v := range e.FieldName {
		if k < len(e.FieldRead) && e.FieldRead[k] != "" {
			javaClass.Fs = append(javaClass.Fs, &JField{
				Name: v,
				Type: e.FieldType[k],
				Mark: e.FieldMark[k],
			})
		}
	}
	return []byte(javaClass.Build())
}

type kv map[string]interface{}

func (e *Excel) ToJson() []byte {
	var arr []kv
	for _, row := range e.rows {
		m := make(kv)
		for i, cell := range row {
			if i < len(e.FieldRead) && e.FieldRead[i] != "" {
				m[e.FieldName[i]] = e.converType(e.FieldType[i], cell)
			}
		}
		arr = append(arr, m)
	}
	jsonBytes, _ := json.Marshal(arr)
	return jsonBytes
}

func (e *Excel) ToProtobuf() []byte {
	message := &Message{
		Name: e.SheetName,
	}
	for k, v := range e.FieldName {
		if k < len(e.FieldRead) && e.FieldRead[k] != "" {
			message.Fs = append(message.Fs, &PField{
				Name: v,
				Type: e.FieldType[k],
				Mark: e.FieldMark[k],
			})
		}
	}
	return []byte(message.Build())
}

func (e *Excel) ToLocal() []byte {
	l := &Local{
		Name: e.SheetName,
	}
	for k, v := range e.FieldName {
		if k < len(e.FieldRead) && e.FieldRead[k] != "" {
			l.Fs = append(l.Fs, &GField{
				Name: v,
				Type: e.FieldType[k],
				Mark: e.FieldMark[k],
			})
		}
	}
	return []byte(l.Build())
}

func (e *Excel) parseRows(rows [][]string) {
	for k, v := range rows {
		if k > FieldName {
			break
		}
		switch k {
		case FieldMark:
			e.FieldMark = v
			break
		case FieldRead:
			e.FieldRead = v
			break
		case FieldType:
			e.FieldType = v
			break
		case FieldName:
			e.FieldName = v
			break
		}
	}
	e.rows = rows[FieldName+1:]
}

func (e *Excel) converType(t string, v string) interface{} {
	var val interface{}
	switch t {
	case "int":
		val, _ = strconv.Atoi(v)
		break
	case "float":
		val, _ = strconv.ParseFloat(v, 32)
		break
	case "lang":
		val = map[string]string{"key": v}
		break
	default:
		val = v
	}
	return val
}

func (e *Excel) setSheetName(name string) {
	i := strings.LastIndex(name, "#")
	if i >= 0 {
		name = name[:i]
	}
	e.SheetName = name
}
