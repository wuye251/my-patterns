package my_prototype

import "fmt"

type File struct {
	Name string
}

func (f *File) Print(input string) {
	fmt.Println(input + f.Name)
}

func (f *File) Clone() Node {
	return &File{Name: f.Name + "_clone"}
}
