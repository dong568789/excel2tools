package main

import (
	"github.com/dong568789/excel2tools/lib"
)

var (
	output = "data"
	input  = "D-道具表.xlsx"
)

func main() {
	e := lib.NewExcel(output)
	e.Read(input)
	e.Export()
}
