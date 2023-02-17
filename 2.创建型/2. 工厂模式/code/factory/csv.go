package factory

type Csv struct {
	Export
}

func newCsv(name, suffix string) *Csv {
	return &Csv{
		Export: Export{
			name:   name,
			suffix: suffix,
		},
	}
}
