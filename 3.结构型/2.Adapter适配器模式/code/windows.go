package adapter

import "fmt"

type Windows struct {
}

func (win *Windows) InsertIntoWindowsLightningPort() {
	fmt.Println("Windows InsertIntoLightningPort()...")
}
