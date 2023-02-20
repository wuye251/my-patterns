package my_builder

type Builder interface {
	BuildHead()
	BuildBody()
	BuildArmLeft()
	BuildArmRight()
	BuildLegLeft()
	BuildLegRight()
	// GetPerson() Builder
}

func NewPersonInstance(flag string) Builder {
	if flag == "fat" {
		return &PersonFatBuilder{}
	}
	if flag == "thin" {
		return &PersonThinBuilder{}
	}

	return nil
}
