package directions

import (
	motors "mqtt-robo-car-v1/robo-car-actions/directions/motors"
	systemconfigreader "mqtt-robo-car-v1/systemconfigreader"
)

var config systemconfigreader.MotorsConfig = systemconfigreader.GetMotorsConfigParams()

/*
const (
	motorAIn1GpioPin    = 2
	motorAIn2GpioPin    = 3
	motorAPWMChannel    = "pwm0"
	motorAPWMController = "pwmchip0"

	motorBIn1GpioPin    = 4
	motorBIn2GpioPin    = 17
	motorBPWMChannel    = "pwm1"
	motorBPWMController = "pwmchip0"
)
*/

func Motors() (motors.Motor, motors.Motor) {
	MotA := motors.NewMotor(config.MotorAIn1GpioPin, config.MotorAIn2GpioPin, config.MotorAPWMController, config.MotorAPWMChannel)
	MotB := motors.NewMotor(config.MotorBIn1GpioPin, config.MotorBIn2GpioPin, config.MotorBPWMController, config.MotorBPWMChannel)
	return MotA, MotB
}

func Foward(MotA motors.Motor, MotB motors.Motor) {
	MotA.Foward()
	MotB.Foward()
}

func Left(MotA motors.Motor, MotB motors.Motor) {
	MotA.Backward()
	MotB.Foward()
}

func Right(MotA motors.Motor, MotB motors.Motor) {
	MotA.Foward()
	MotB.Backward()
}

func Stop(MotA motors.Motor, MotB motors.Motor) {
	MotA.Stop()
	MotB.Stop()
}

func SetSpeed(MotA motors.Motor, MotB motors.Motor, dutyCycle int) {
	MotA.SetSpeed(dutyCycle)
	MotB.SetSpeed(dutyCycle)
}
