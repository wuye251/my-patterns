package factory

type Export struct {
	name   string
	suffix string
}

func (e *Export) setName(name string) {
	e.name = name
}

func (e *Export) setSuffix(suffix string) {
	e.suffix = suffix
}

func (e *Export) GetName() string {
	return e.name
}

func (e *Export) GetSuffix() string {
	return e.suffix
}
