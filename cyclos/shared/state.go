package shared


const (
	Home     = UIPage("home")
	Stat     = UIPage("stat")
	Nav      = UIPage("nav")
	Trends   = UIPage("trends")
	Settings = UIPage("settings")

	KmH = SpeedUnit("km/h")

	NetNotAvailable = NetworkStrength(-1)
	Disconnected = NetworkStrength(0)
	Low = NetworkStrength(1)
	Medium = NetworkStrength(2)
	High = NetworkStrength(3)
	Full = NetworkStrength(4)

	BatNotAvailable = BatteryStatus(-1)
	Charging = BatteryStatus(0)
	Error = BatteryStatus(1)
	ThirdsPrecision = BatteryStatus(2)
	PercentagePrecision = BatteryStatus(3)
)

type UIPage string

//InputState provides state controlled by user input.
type InputState struct {
	UIPage        UIPage
}

//Clock provides the current time, if available.
type Clock struct {
	Available bool
	Current   string
}

type SpeedUnit string

//RTMetrics provides live processed metrics pulled from metrics hardware. RT stands for real-time. RTMetrics is provded by the RTDATA subsystem.
type RTMetrics struct {
	Speed int //in selected unit
	SpeedUnit SpeedUnit
}

type NetworkStrength int

//Misc provides miscellaneous data that, in general, is cheap to gather and does not require much processing. Misc is provided by the GENERALDATA.
type Misc struct {
	Clock Clock
	Battery Battery
	NetworkStrength NetworkStrength
	DisplayWidth  int16
	DisplayHeight int16
}

type BatteryStatus int

//Battery provides battery charge information, if available.
type Battery struct {
	Status BatteryStatus
	Value float64 //a decimal value representing battery charge, from 0 (fully discharged) to 1 (fully charged). Precision may vary depending on hardware and BatteryStatus.
}

//Statistics provides over-time processed statistics. Statistics is, Unlike RTMetrics, managed by the NAVSTAT subsystem.
type Statistics struct {

}

type SystemState struct {
	Input InputState
	Misc Misc
	RTMetrics RTMetrics
	Statistics Statistics
}
