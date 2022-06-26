package global

import (
	"server/pkg/setting"
)

var (
	ServerSettings   *setting.ServerSettings
	AppSettings      *setting.AppSettings
	DatabaseSettings *setting.DatabaseSettings
)

func SetupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSettings("Server", &ServerSettings)
	if err != nil {
		return err
	}

	err = setting.ReadSettings("App", &AppSettings)
	if err != nil {
		return err
	}

	err = setting.ReadSettings("Database", &DatabaseSettings)
	if err != nil {
		return err
	}

	return nil
}
