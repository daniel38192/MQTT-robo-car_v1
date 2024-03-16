package motors

import (
	linuxgpio "github.com/daniel38192/go-gpio/gpio"
	"github.com/daniel38192/go-gpio/gpio/enums"
	linuxpwm "github.com/daniel38192/go-gpio/pwm"
	"github.com/daniel38192/go-gpio/pwm/polarity"
)

type Motor struct {
	motorInput1GpioPin linuxgpio.GPIO
	motorInput2GpioPin linuxgpio.GPIO
	motorEnablePWM     linuxpwm.PWM
}

func NewMotor(motorInput1GpioPinINT int, motorInput2GpioPinINT int, motorEnablePWMController string, motorEnablePWMChannel string) Motor {
	motorInput1GpioPin := linuxgpio.NewGpio(motorInput1GpioPinINT, false, enums.OUT)
	motorIn2GpioPin := linuxgpio.NewGpio(motorInput2GpioPinINT, false, enums.OUT)
	motorEnablePWM := linuxpwm.NewPWM(motorEnablePWMController, motorEnablePWMChannel, 1000000, polarity.Normal)
	// Max duty cycle possible "500000" for 50%
	motorEnablePWM.SetDutyCycle(1000000)
	motorEnablePWM.Enable(true)
	motor := Motor{motorInput1GpioPin: motorInput1GpioPin, motorInput2GpioPin: motorIn2GpioPin, motorEnablePWM: motorEnablePWM}
	return motor
}

func (motor Motor) Foward() {
	motor.motorInput1GpioPin.WriteGpioValue(true)
	motor.motorInput2GpioPin.WriteGpioValue(false)
}

func (motor Motor) Backward() {
	motor.motorInput1GpioPin.WriteGpioValue(false)
	motor.motorInput2GpioPin.WriteGpioValue(true)
}

func (motor Motor) Stop() {
	motor.motorInput1GpioPin.WriteGpioValue(false)
	motor.motorInput2GpioPin.WriteGpioValue(false)
}

func (motor Motor) SetSpeed(dutyCycle int) {
	motor.motorEnablePWM.SetDutyCycle(dutyCycle)
}
