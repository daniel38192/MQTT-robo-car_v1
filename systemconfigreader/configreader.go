package systemconfigreader

import (
	"fmt"
	defaultconfig "mqtt-robo-car-v1/systemconfigreader/defaultconfig"
	"os"

	viper "github.com/spf13/viper"
)

type ClientConfig struct {
	Ip                  string
	Port                string
	ClientID            string
	AnonymousConnection bool
	DirectionTopic      string
	SpeedTopic          string
	Username            string
	Password            string
}

type RoboCarConfig struct {
	FowardMessage          string
	LeftMessage            string
	RightMessage           string
	StopMessage            string
	AcknowledgementTopic   string
	AcknowledgementMessage string
}

type MotorsConfig struct {
	MotorAIn1GpioPin    int
	MotorAIn2GpioPin    int
	MotorAPWMChannel    string
	MotorAPWMController string

	MotorBIn1GpioPin    int
	MotorBIn2GpioPin    int
	MotorBPWMChannel    string
	MotorBPWMController string
}

func GetClientConfigParams() ClientConfig {
	loadConfig()
	clientConfig := ClientConfig{
		Ip:                  viper.GetString("broker.ip"),
		Port:                viper.GetString("broker.port"),
		ClientID:            viper.GetString("broker.client_id"),
		AnonymousConnection: viper.GetBool("broker.anonymous_connection"),
		DirectionTopic:      viper.GetString("topics.direction_topic"),
		SpeedTopic:          viper.GetString("topics.speed_topic"),
		Username:            viper.GetString("broker.username"),
		Password:            viper.GetString("broker.password"),
	}
	return clientConfig
}

func GetRoboCarConfigParams() RoboCarConfig {
	loadConfig()
	roboCarConfig := RoboCarConfig{
		FowardMessage:          viper.GetString("messages.forward_message"),
		LeftMessage:            viper.GetString("messages.left_message"),
		RightMessage:           viper.GetString("messages.right_message"),
		StopMessage:            viper.GetString("messages.stop_message"),
		AcknowledgementTopic:   viper.GetString("topics.acknowledgement_topic"),
		AcknowledgementMessage: viper.GetString("messages.acknowledgement_message"),
	}
	return roboCarConfig
}

func GetMotorsConfigParams() MotorsConfig {
	motorsConfig := MotorsConfig{
		MotorAIn1GpioPin:    viper.GetInt("gpios.motor_a_in_1_gpio_pin"),
		MotorAIn2GpioPin:    viper.GetInt("gpios.motor_a_in_2_gpio_pin"),
		MotorAPWMChannel:    viper.GetString("gpios.motor_a_pwm_channel"),
		MotorAPWMController: viper.GetString("gpios.motor_a_pwm_controller"),

		MotorBIn1GpioPin:    viper.GetInt("gpios.motor_b_in_1_gpio_pin"),
		MotorBIn2GpioPin:    viper.GetInt("gpios.motor_b_in_2_gpio_pin"),
		MotorBPWMChannel:    viper.GetString("gpios.motor_b_pwm_channel"),
		MotorBPWMController: viper.GetString("gpios.motor_b_pwm_controller"),
	}
	return motorsConfig
}

func loadConfig() {
	viper.AddConfigPath(defaultconfig.LinuxSysPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		loadLocalConfig()
	}
}

func loadLocalConfig() {
	viper.AddConfigPath(defaultconfig.LocalPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		createExampleConfig(err)
	}
}

func createExampleConfig(err error) {
	fmt.Println("Config file not found!")
	os.WriteFile(defaultconfig.LocalPath+"/config.yml", []byte(defaultconfig.Defaultconfig), 0700)
	panic(err)
}
