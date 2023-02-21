package my_template

import "fmt"

type PdfExport struct {
	Export
	Data [][]interface{}
}

func (pdf *PdfExport) AnalyzeData(data [][]interface{}) {
	fmt.Printf("analyze pdf data\n")
}

func (pdf *PdfExport) WriteDataToFile() {
	fmt.Printf("write pdf data to file\n")
}
