package core

import (
	"fmt"
	"runtime"
	"time"
)

const (
	appConfigError      = "App configuration error: %s"
	defaultPingInterval = 1 * time.Minute
	defaultSendInterval = 500 * time.Millisecond
)

type Configuration struct {
	Id            int32
	Hash          string
	Version       string
	DeviceModel   string
	SystemVersion string
	Language      string
	//SessionHome   string
	PingInterval time.Duration
	SendInterval time.Duration
	KeyPath      string
}

func NewConfiguration(id int32, hash, version, deviceModel, systemVersion, language string, pingInterval time.Duration, sendInterval time.Duration, keyPath string) (Configuration, error) {
	//appConfig := new(Configuration)
	appConfig := Configuration{}

	if id == 0 || hash == "" || version == "" {
		return Configuration{}, fmt.Errorf(appConfigError, "Fields Id, Hash or Version are empty")
	}
	appConfig.Id = id
	appConfig.Hash = hash
	appConfig.Version = version

	appConfig.DeviceModel = deviceModel
	if deviceModel == "" {
		appConfig.DeviceModel = "Unknown"
	}

	appConfig.SystemVersion = systemVersion
	if systemVersion == "" {
		appConfig.SystemVersion = runtime.GOOS + "/" + runtime.GOARCH
	}

	appConfig.Language = language
	if language == "" {
		appConfig.Language = "en"
	}

	//appConfig.SessionHome = sessionFileHome
	//if sessionFileHome == "" {
	//	usr, err := user.Current()
	//	if err != nil {
	//		appConfig.SessionHome = os.Getenv("HOME")
	//	} else {
	//		appConfig.SessionHome = usr.HomeDir
	//	}
	//}
	appConfig.KeyPath = keyPath

	appConfig.PingInterval = pingInterval
	if pingInterval == 0 {
		appConfig.PingInterval = defaultPingInterval
	}

	appConfig.SendInterval = sendInterval
	if sendInterval == 0 {
		appConfig.SendInterval = defaultSendInterval
	}

	return appConfig, nil
}

func (appConfig Configuration) Check() error {
	if appConfig.Id == 0 || appConfig.Hash == "" || appConfig.Version == "" {
		return fmt.Errorf(appConfigError, "Configuration.Id, Configuration.Hash or Configuration.Version are empty")
	}

	if appConfig.DeviceModel == "" {
		return fmt.Errorf(appConfigError, "Configuration.DeviceModel is empty")
	}

	if appConfig.SystemVersion == "" {
		return fmt.Errorf(appConfigError, "Configuration.SystemVersion is empty")
	}

	if appConfig.Language == "" {
		return fmt.Errorf(appConfigError, "Configuration.Language is empty")
	}

	return nil
}
