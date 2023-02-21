package my_template

import "fmt"

type IExport interface {
	AnalyzeData([][]interface{})
	WriteDataToFile()
	Upload(path string)
}

// 基类实现公共方法
type Export struct {
	Iexport IExport
}

func (export *Export) Upload(path string) {
	fmt.Printf("base export upload file to remote server, path is %s\n", path)
}
