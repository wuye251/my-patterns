package main

type Computer struct{}

func (cg *Computer) Start() {
	println("Computer start")
}

func (cg *Computer) Stop() {
	println("Computer stop")
}

type TV struct{}

func (tv *TV) Open() {
	println("TV open")
}

func (tv *TV) Close() {
	println("TV close")
}

type WashingMachine struct{}

func (wm *WashingMachine) Run() {
	println("WashingMachine run")
}

func (wm *WashingMachine) Stop() {
	println("WashingMachine stop")
}

type Facade struct {
	cg *Computer
	tv *TV
	wm *WashingMachine
}

// 对外的接口打开一套设备
func (f *Facade) Play() {
	f.cg.Start()
	f.tv.Open()
	f.wm.Run()
}

// 对外的接口关闭一套设备
func (f *Facade) Stop() {
	f.cg.Stop()
	f.tv.Close()
	f.wm.Stop()
}

func main() {
	f := Facade{&Computer{}, &TV{}, &WashingMachine{}}
	f.Play()
	f.Stop()
}
