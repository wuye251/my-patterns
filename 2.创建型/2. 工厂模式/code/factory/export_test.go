package factory_test

import (
	"fmt"
	factory "my-factory/factory"
	"testing"
)

// command: `go test -v -timeout 30s -run ^TestExport$ my-factory/factory`
func TestExport(t *testing.T) {
	setSuffix := "csv"
	setName := "csvName"
	instance, err := factory.GetExportInstance(setName, setSuffix)
	if err != nil {
		t.Fatalf("err is %s\n", err.Error())
	}
	fmt.Printf("name is %s\n", instance.GetName())
	fmt.Printf("suffix is %s\n", instance.GetSuffix())

	setSuffix1 := "xlsx"
	setName1 := "xlsxName"
	instance1, err := factory.GetExportInstance(setName1, setSuffix1)
	if err != nil {
		t.Fatalf("err is %s\n", err.Error())
	}
	fmt.Printf("name is %s\n", instance1.GetName())
	fmt.Printf("suffix is %s\n", instance1.GetSuffix())
}
