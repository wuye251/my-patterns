package my_template_test

import (
	template "my-template"
	"testing"
)

// command: `go test -v -timeout 30s -run ^TestExprot$ my-template`
func TestExprot(t *testing.T) {
	excel := &template.ExcelExport{}
	excel.Export = template.Export{
		Iexport: excel,
	}
	excel.AnalyzeData([][]interface{}{})
	excel.WriteDataToFile()
	excel.Upload("excel file path")

	pdf := &template.PdfExport{}
	pdf.Export = template.Export{
		Iexport: pdf,
	}
	pdf.AnalyzeData([][]interface{}{})
	pdf.WriteDataToFile()
	pdf.Upload("pdf file path")
}
