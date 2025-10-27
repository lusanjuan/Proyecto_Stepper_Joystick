package joystick

import "machine"

type Joystick struct {
	xPin  machine.ADC
	yPin  machine.ADC
	swPin machine.Pin
}

// Crear nuevo joystick
func NewJoystick(x machine.Pin, y machine.Pin, sw machine.Pin) *Joystick {
	j := &Joystick{
		xPin:  machine.ADC{Pin: x},
		yPin:  machine.ADC{Pin: y},
		swPin: sw,
	}
	j.xPin.Configure(machine.ADCConfig{})
	j.yPin.Configure(machine.ADCConfig{})
	j.swPin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	return j
}

// Normalizado de -1 a +1
func (j *Joystick) ReadAxisX() float32 {
	raw := j.xPin.Get()
	x := (float32(raw)/65535.0 - 0.5) * 2
	return x
}

func (j *Joystick) ReadAxisY() float32 {
	raw := j.yPin.Get()
	y := (float32(raw)/65535.0 - 0.5) * 2
	return y
}

// True si el botón está presionado
func (j *Joystick) ButtonPressed() bool {
	return !j.swPin.Get() // Pull-up: LOW = presionado
}
