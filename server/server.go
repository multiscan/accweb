package server

import (
	"os/exec"
)

type ServerSettings struct {
	Id  int       `json:"id"`
	PID int       `json:"pid"` // 0 = stopped, else running
	Cmd *exec.Cmd `json:"-"`

	// ACC server configuration files
	Configuration ConfigurationJson `json:"basic"`
	Settings      SettingsJson      `json:"settings"`
	Event         EventJson         `json:"event"`
	EventRules    EventRulesJson    `json:"eventRules"`
	Entrylist     EntrylistJson     `json:"entrylist"`
}

type ConfigurationJson struct {
	ConfigVersion   int `json:"configVersion"`
	UdpPort         int `json:"udpPort"`
	TcpPort         int `json:"tcpPort"`
	MaxConnections  int `json:"maxConnections"`
	RegisterToLobby int `json:"registerToLobby"`
	LanDiscovery    int `json:"lanDiscovery"`
}

type SettingsJson struct {
	ConfigVersion              int    `json:"configVersion"`
	ServerName                 string `json:"serverName"`
	Password                   string `json:"password"`
	AdminPassword              string `json:"adminPassword"`
	TrackMedalsRequirement     int    `json:"trackMedalsRequirement"`
	SafetyRatingRequirement    int    `json:"safetyRatingRequirement"`
	RacecraftRatingRequirement int    `json:"racecraftRatingRequirement"`
	SpectatorPassword          string `json:"spectatorPassword"`
	DumpLeaderboards           int    `json:"dumpLeaderboards"`
	IsRaceLocked               int    `json:"isRaceLocked"`
	RandomizeTrackWhenEmpty    int    `json:"randomizeTrackWhenEmpty"`
	MaxCarSlots                int    `json:"maxCarSlots"`
	CentralEntryListPath       string `json:"centralEntryListPath"`
	ShortFormationLap          int    `json:"shortFormationLap"`
	AllowAutoDQ                int    `json:"allowAutoDQ"`
	DumpEntryList              int    `json:"dumpEntryList"`
}

type EventJson struct {
	ConfigVersion             int               `json:"configVersion"`
	Track                     string            `json:"track"`
	EventType                 string            `json:"eventType"`
	PreRaceWaitingTimeSeconds int               `json:"preRaceWaitingTimeSeconds"`
	SessionOverTimeSeconds    int               `json:"sessionOverTimeSeconds"`
	AmbientTemp               int               `json:"ambientTemp"`
	CloudLevel                float64           `json:"cloudLevel"`
	Rain                      float64           `json:"rain"`
	WeatherRandomness         int               `json:"weatherRandomness"`
	Sessions                  []SessionSettings `json:"sessions"`
	PostQualySeconds          int               `json:"postQualySeconds"`
	PostRaceSeconds           int               `json:"postRaceSeconds"`
}

type EventRulesJson struct {
	ConfigVersion                        int  `json:"configVersion"`
	QualifyStandingType                  int  `json:"qualifyStandingType"`
	PitWindowLengthSec                   int  `json:"pitWindowLengthSec"`
	DriverStintTimeSec                   int  `json:"driverStintTimeSec"`
	MandatoryPitstopCount                int  `json:"mandatoryPitstopCount"`
	MaxTotalDrivingTime                  int  `json:"maxTotalDrivingTime"`
	MaxDriversCount                      int  `json:"maxDriversCount"`
	IsRefuellingAllowedInRace            bool `json:"isRefuellingAllowedInRace"`
	IsRefuellingTimeFixed                bool `json:"isRefuellingTimeFixed"`
	IsMandatoryPitstopRefuellingRequired bool `json:"isMandatoryPitstopRefuellingRequired"`
	IsMandatoryPitstopTyreChangeRequired bool `json:"isMandatoryPitstopTyreChangeRequired"`
	IsMandatoryPitstopSwapDriverRequired bool `json:"isMandatoryPitstopSwapDriverRequired"`
}

type SessionSettings struct {
	HourOfDay              int    `json:"hourOfDay"`
	DayOfWeekend           int    `json:"dayOfWeekend"`
	TimeMultiplier         int    `json:"timeMultiplier"`
	SessionType            string `json:"sessionType"`
	SessionDurationMinutes int    `json:"sessionDurationMinutes"`
}

type EntrylistJson struct {
	ConfigVersion  int             `json:"configVersion"`
	Entries        []EntrySettings `json:"entries"`
	ForceEntryList int             `json:"forceEntryList"`
}

type EntrySettings struct {
	Drivers                      []DriverSettings `json:"drivers"`
	RaceNumber                   int              `json:"raceNumber"`
	ForcedCarModel               int              `json:"forcedCarModel"`
	OverrideDriverInfo           int              `json:"overrideDriverInfo"`
	IsServerAdmin                int              `json:"isServerAdmin"`
	CustomCar                    string           `json:"customCar"`
	OverrideCarModelForCustomCar int              `json:"overrideCarModelForCustomCar"`
}

type DriverSettings struct {
	FirstName           string `json:"firstName"`
	LastName            string `json:"lastName"`
	ShortName           string `json:"shortName"`
	DriverCategory      int    `json:"driverCategory"`
	PlayerID            string `json:"playerID"`
	DefaultGridPosition int    `json:"defaultGridPosition"`
}

func (server *ServerSettings) start(cmd *exec.Cmd) {
	server.PID = cmd.Process.Pid
	server.Cmd = cmd
}

func (server *ServerSettings) stop() {
	server.PID = 0
	server.Cmd = nil
}
