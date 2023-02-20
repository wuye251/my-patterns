package my_prototype

import "fmt"

type Folder struct {
	Children []Node
	Name     string
}

func (f *Folder) Print(input string) {
	fmt.Println(input + f.Name)
	for _, val := range f.Children {
		val.Print(input + input)
	}
}

func (f *Folder) Clone() Node {
	clonefolder := &Folder{Name: f.Name + "_clone"}
	var tmpChildren []Node
	for _, i := range f.Children {
		copy := i.Clone()
		tmpChildren = append(tmpChildren, copy)
	}
	clonefolder.Children = tmpChildren

	return clonefolder
}
