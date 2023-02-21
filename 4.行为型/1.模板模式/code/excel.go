package my_template

import "fmt"

type ExcelExport struct {
	Export
	Data [][]interface{}
}

func (excel *ExcelExport) AnalyzeData(data [][]interface{}) {
	fmt.Printf("analyze excel data\n")
}

func (excel *ExcelExport) WriteDataToFile() {
	fmt.Printf("write excel data to file\n")
}
