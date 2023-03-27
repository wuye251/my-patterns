package adapter_test

import (
	"adapter"
	"testing"
)

func TestAdapter(t *testing.T) {
	client := adapter.Client{}

	mac := &adapter.Mac{}
	
	client.InsertIntoLightningConnectIntoComputer(mac)
	// Error
	// win := &adapter.Windows{}
	// client.InsertIntoLightningConnectIntoComputer(win)

	// add win adapter
	winAdapter := &adapter.WindowsAdapter{}
	client.InsertIntoLightningConnectIntoComputer(winAdapter)

}
