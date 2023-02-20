package my_builder

type Director struct {
	builder Builder
}

func NewDirector() *Director {
	return &Director{}
}

func (dir *Director) SetBuilder(b Builder) {
	dir.builder = b
}

func (dir *Director) Build() {
	dir.builder.BuildHead()
	dir.builder.BuildBody()
	dir.builder.BuildArmLeft()
	dir.builder.BuildArmRight()
	dir.builder.BuildLegLeft()
	dir.builder.BuildLegRight()
}
