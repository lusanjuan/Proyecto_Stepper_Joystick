package main

import (
	"machine"
	"time"
	"fmt"
	"proyecto_stepper_joystick/joystick"
	"proyecto_stepper_joystick/stepper"
)

func main() {
	machine.InitADC()

	js := joystick.NewJoystick(machine.ADC0, machine.ADC1, machine.GP15)
	motor := stepper.NewStepper(machine.GP10, machine.GP11, machine.GP17)

	fmt.Println("Inicio :)")

	for {
		x := js.ReadAxisX()
		button := js.ButtonPressed()

		if button {
			motor.Stop()
			fmt.Println("Botón presionado -> Motor detenido")
			time.Sleep(500 * time.Millisecond)
			continue
		}

		if x > 0.1 {
			motor.Enable()
			motor.Step(1)
			delay := time.Duration(5 + (1-x)*15) 
			fmt.Printf("Giro hacia derecha | X=%.2f | delay=%dms\n", x, delay)
			time.Sleep(delay * time.Millisecond)

		} else if x < -0.1 {
			motor.Enable()
			motor.Step(-1)
			delay := time.Duration(5 + (1-(-x))*15) 
			fmt.Printf("Giro hacia izquierda | X=%.2f | delay=%dms\n", x, delay)
			time.Sleep(delay * time.Millisecond)

		} else {
			motor.Stop()
			fmt.Println("Joystick en reposo")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
