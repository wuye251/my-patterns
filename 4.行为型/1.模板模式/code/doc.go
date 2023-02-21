package my_template

import "fmt"

type DocExport struct {
	Export
	Data [][]interface{}
}

func (doc *DocExport) AnalyzeData(data [][]interface{}) {
	fmt.Printf("analyze doc data\n")
}

func (doc *DocExport) WriteDataToFile() {
	fmt.Printf("write doc data to file\n")
}
