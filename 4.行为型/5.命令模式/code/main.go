package main

import "fmt"

type Doctor struct{}

func (d *Doctor) TreatNose() {
	fmt.Println("治疗鼻子")
}

func (d *Doctor) TreadEye() {
	fmt.Println("治疗眼睛")
}

// 对外的接口
type Commond interface {
	Treat()
}

// 对外的治疗鼻子接口
type CommondTreatNose struct {
	doctor *Doctor
}

func (cmdTreatNose *CommondTreatNose) Treat() {
	cmdTreatNose.doctor.TreatNose()
}

// 对外的治疗眼睛接口
type CommondTreatEye struct {
	doctor *Doctor
}

func (cmdTreatEye *CommondTreatEye) Treat() {
	cmdTreatEye.doctor.TreadEye()
}

// 护士
type Nurse struct {
	Cmd []Commond
}

func (nurse *Nurse) Recive() {
	if nurse.Cmd == nil {
		return
	}

	for _, cmd := range nurse.Cmd {
		cmd.Treat()
	}
}

func main() {
	// 找护士
	nurse := new(Nurse)
	nurse.Cmd = append(nurse.Cmd, new(CommondTreatEye))
	nurse.Cmd = append(nurse.Cmd, new(CommondTreatNose))

	nurse.Recive()
}
