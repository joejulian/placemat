package ipmi

// port from OpenIPMI
// Sensor/Event Network Function
const (
	IPMI_CMD_SET_EVENT_RECEIVER =			0x00
	IPMI_CMD_GET_EVENT_RECEIVER =			0x01
	IPMI_CMD_PLATFORM_EVENT =			0x02

	IPMI_CMD_GET_PEF_CAPABILITIES =			0x10
	IPMI_CMD_ARM_PEF_POSTPONE_TIMER =		0x11
	IPMI_CMD_SET_PEF_CONFIG_PARMS =			0x12
	IPMI_CMD_GET_PEF_CONFIG_PARMS =			0x13
	IPMI_CMD_SET_LAST_PROCESSED_EVENT_ID =		0x14
	IPMI_CMD_GET_LAST_PROCESSED_EVENT_ID =		0x15
	IPMI_CMD_ALERT_IMMEDIATE =			0x16
	IPMI_CMD_PET_ACKNOWLEDGE =			0x17

	IPMI_CMD_GET_DEVICE_SDR_INFO =			0x20
	IPMI_CMD_GET_DEVICE_SDR =			0x21
	IPMI_CMD_RESERVE_DEVICE_SDR_REPOSITORY =	0x22
	IPMI_CMD_GET_SENSOR_READING_FACTORS =		0x23
	IPMI_CMD_SET_SENSOR_HYSTERESIS =		0x24
	IPMI_CMD_GET_SENSOR_HYSTERESIS =		0x25
	IPMI_CMD_SET_SENSOR_THRESHOLD =			0x26
	IPMI_CMD_GET_SENSOR_THRESHOLD =			0x27
	IPMI_CMD_SET_SENSOR_EVENT_ENABLE =		0x28
	IPMI_CMD_GET_SENSOR_EVENT_ENABLE =		0x29
	IPMI_CMD_REARM_SENSOR_EVENTS =			0x2a
	IPMI_CMD_GET_SENSOR_EVENT_STATUS =		0x2b
	IPMI_CMD_GET_SENSOR_READING =			0x2d
	IPMI_CMD_SET_SENSOR_TYPE =			0x2e
	IPMI_CMD_GET_SENSOR_TYPE =			0x2f
)
