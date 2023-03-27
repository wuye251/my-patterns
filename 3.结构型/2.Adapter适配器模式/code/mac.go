package adapter

import "fmt"

type Mac struct {
}

func (mac *Mac) InsertIntoLightningPort() {
	fmt.Println("mac InsertIntoLightningPort()...")
}
