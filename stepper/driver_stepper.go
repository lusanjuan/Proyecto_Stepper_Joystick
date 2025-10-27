package stepper

import (
	"machine"
)

type Stepper struct {
	step   machine.Pin
	dir    machine.Pin
	enable machine.Pin
}


func NewStepper(stepPin, dirPin, enablePin machine.Pin) *Stepper {
	s := &Stepper{
		step:   stepPin,
		dir:    dirPin,
		enable: enablePin,
	}

	s.step.Configure(machine.PinConfig{Mode: machine.PinOutput})
	s.dir.Configure(machine.PinConfig{Mode: machine.PinOutput})
	s.enable.Configure(machine.PinConfig{Mode: machine.PinOutput})

	s.enable.Low() 
	return s
}

func (s *Stepper) Step(dir int) {
	if dir > 0 {
		s.dir.Low() 
	} else {
		s.dir.High() 
	}
	s.step.High()
	s.step.Low()
}


func (s *Stepper) Stop() {
	s.enable.High() 
}


func (s *Stepper) Enable() {
	s.enable.Low() 
}
