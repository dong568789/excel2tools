package main

import (
	"github.com/dong568789/excel2tools/lib"
	"github.com/dong568789/excel2tools/utils"
	"os"
)

var (
	output = "data"
	input  = ""
)

func main() {
	args := os.Args
	for k, v := range args {
		if k == 1 {
			input = v
		}
	}

	if input == "" {
		utils.Log().Error("请输入文件路径")
		return
	}

	e := lib.NewExcel(output)
	e.Read(input)
	e.Export()
}
