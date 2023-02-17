package factory

import "bytes"

type Xlsx struct {
	filename string
	suffix   string
}

func newXlsx(name, suffix string) *Xlsx {
	return &Xlsx{}
}

func (xlsx *Xlsx) setName(name string) {
	xlsx.filename = name
}
func (xlsx *Xlsx) setSuffix(suffix string) {
	xlsx.suffix = suffix
}
func (xlsx *Xlsx) GetSuffix() string {
	return xlsx.suffix
}
func (xlsx *Xlsx) GetName() string {
	return xlsx.suffix
}

// 写数据到buffer
func (xlsx *Xlsx) WriteDataToBuffer(data [][]interface{}) bytes.Buffer {
	// gen data and write to buffer
	return bytes.Buffer{}
}
