# Proyect: Stepper motor control with joystick

## Description  
This proyect allows controlling the direction and speed of a stepper motor using an analog joystick.
It is programmed in Go using TinyGo for the Raspberry Pi Pico microcontroller.

The system reads the analog values from the joysticks's X and Y axes. In this case, depending on the X-axis position: 

-The rotation direction (left/right) is controlled.

-The rotation speed of the motor is adjusted. 

Aditionally, pressing the joystick's button stops the motor's movement. 

### Hardware used

| Component        | Description       |
|--------------------|---------------------|
| Microcontroller   | Raspberry Pi Pico   |
| Analog joystick | 2 axes + 1 button    |
| Stepper Motor  | 17HS4401            |
| Motor driver   | TMC2209             |


## Wiring

### Joystick

| Component  | Pico (GPIO) |        Description            |
|-------------|-------------|------------------------------|
| X Axis       | ADC0 (GP26) | Analog imput for X axis |
|  Y Axis      | ADC1 (GP27) | Analog imput for Y axis |
| Sw          | GP15        | Digital imput              |
| VCC         | 3.3 V       | Power supply for joystick  |
| GND         | GND         | Common ground                 | 

### Stepper Motor (with  TMC2209 driver)

| TMC2209 Pin| Pico (GPIO) | Motor Coil |
|------------------|-------------|------------------|
| 1A               |             | BLK (Coil A)   |
| 1B               |             | GRN (Coil A)   |
| 2A               |             | RED (Coil B)   |
| 2B               |             | BLU (Coil B)   |
| VDD              | 3V3(OUT)    |                  |
| VMOT             | VBUS        |                  |
| DIR              | GP11        |                  |
| STEP             | GP10        |                  |
| ENABLE           | GP17        |                  |
| GND              | GND         |                  |


## Operation

1. The program initializes the ADCs to read the joystick axis values.
2. The motor’s rotation direction changes according to the position of the joystick along the X-axis (positive/negative).
3. The speed rotation is calculated following:

```bash
v = 5 + (1-x)*15
```

where "x" si the normalized X-axis value in the range [-1,1]
4. When the joystick button is pressed, the motor stops (ENABLE pin deactivated)

## Compile and upload the program

### 1. Install TinyGo
Download TinyGo from: 

[https://tinygo.org/getting-started/](https://tinygo.org/getting-started/)

Verify with:

```bash
tinygo version
```

### 2. Conect the Raspberry Pi Pico in Bootloader mode


1. Hold down the Pico's BOOTSEL button  
2. Connect via USB 
3. Release the BOOTSEL button
4. A new drive name RPI-RP2 should appear 


### 3. Compile and flash the program

From the root directory run: 

```bash
tinygo flash -target=pico main_proyecto.go
```


### 4. View serial output

Open the serial monitor to verify the joysticks behavior:

```bash
tinygo monitor
```


