package factory

type Xlsx struct {
	Export
}

func newXlsx(name, suffix string) *Xlsx {
	return &Xlsx{
		Export: Export{
			name:   name,
			suffix: suffix,
		},
	}
}
