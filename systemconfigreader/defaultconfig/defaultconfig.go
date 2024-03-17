package defaultconfig

type Defconfig string

const (
	Defaultconfig = `
# This is the robo-car yml configuration example.
# These parameters are examples and may be changed.
broker:
# Set the broker ip address, it may be a ip or a dns.
    ip:         "127.0.0.1"
# Set the broker listening port.
    port:       "61616"
# Set the client id 
    client_id:  "robo-car"
# If is true, the robo-car will connect itself as an anonymous user, otherwise specify the user and password below.
    anonymous_connection: false
# Set the username
    username:   "robo-car"
# Also, Set the password for the username given above.
    password:   "secret"
# If is an anonymous login, you can comment the username and password.

topics:
# Set the robo-car direction topic, in this topic the direction messages will be received.
    direction_topic:       "robo-car/Direction"
# Set the robo-car speed topic, in this topic the speed messages will be received.
    speed_topic:            "robo-car/Speed"
# Set the robo-car acknowledgement topic, in this topic the acknowledgement messages will be sended.				
    acknowledgement_topic:  "robo-car/Acknowledgement"

messages:
# Specify the mqtt messages for each action.
    forward_message:        "Forward"
    left_message:           "Left"
    right_message:          "Right"
    stop_message:           "Stop"
# Specify the message when the robo-car has made a action succesfully.
    acknowledgement_message:    "Acknowledgement"

gpios:
# This is a l293d/l298n program based.
# Set In1 gpio pin and In2 gpio pin (integer)
    motor_a_in_1_gpio_pin:    2
    motor_a_in_2_gpio_pin:    3
# Set the linux pwm controller and a channel for exporting it.
    motor_a_pwm_channel:      "pwm0"
    motor_a_pwm_controller:   "pwmchip0"

    motor_b_in_1_gpio_pin:    4
    motor_b_in_2_gpio_pin:    17
    motor_b_pwm_channel:      "pwm1"
    motor_b_pwm_controller:   "pwmchip0"
`
	LinuxSysPath = "/etc/robo-car/"
	LocalPath    = "."
)
