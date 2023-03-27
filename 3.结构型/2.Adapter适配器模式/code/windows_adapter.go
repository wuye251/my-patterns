package adapter

import "fmt"

type WindowsAdapter struct {
}

func (winAdapter *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("winAdapter InsertIntoLightningPort()...")
}
