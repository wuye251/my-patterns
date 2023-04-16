package factory

import "bytes"

type Csv struct {
	filename string
	suffix   string
}

func newCsv(name, suffix string) *Csv {
	return &Csv{}
}

func (csv *Csv) setName(name string) {
	csv.filename = name
}
func (csv *Csv) setSuffix(suffix string) {
	csv.suffix = suffix
}
func (csv *Csv) GetSuffix() string {
	return csv.suffix
}
func (csv *Csv) GetName() string {
	return csv.suffix
}

// 写数据到buffer
func (csv *Csv) WriteDataToBuffer(data [][]interface{}) bytes.Buffer {
	// gen data and write to buffer
	return bytes.Buffer{}
}
