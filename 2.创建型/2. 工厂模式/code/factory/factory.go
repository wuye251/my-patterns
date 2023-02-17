package factory

import "fmt"

func GetExportInstance(name, suffix string) (IExport, error) {
	switch suffix {
	case "csv":
		return newCsv(name, suffix), nil
	case "xlsx":
		return newXlsx(name, suffix), nil
	default:
		return nil, fmt.Errorf("unexpected suffix")
	}
}
