package my_config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	CurrentPatentConfig *PatentConfig
	onceConfig          sync.Once
)

type PatentConfig struct {
	InnojoyUser         string
	InnojoyPassword     string
	FileUrl             string
	NoveltyReportConfig NoveltyReportConfig
}

func LoadPatentConfig() {
	onceConfig.Do(func() {
		CurrentPatentConfig = &PatentConfig{
			InnojoyUser:     viper.GetString("settings.patent.innojoy.user"),
			InnojoyPassword: viper.GetString("settings.patent.innojoy.password"),
			FileUrl:         viper.GetString("settings.files.url"),
			NoveltyReportConfig: NoveltyReportConfig{
				DepartName:  viper.GetString("settings.novelty_report.depart_name"),
				ContactAddr: viper.GetString("settings.novelty_report.contact_addr"),
				ZipCode:     viper.GetString("settings.novelty_report.zip_code"),
				ManagerName: viper.GetString("settings.novelty_report.manager_name"),
				ManagerTel:  viper.GetString("settings.novelty_report.manager_tel"),
				ContactName: viper.GetString("settings.novelty_report.contact_name"),
				ContactTel:  viper.GetString("settings.novelty_report.contact_tel"),
				Email:       viper.GetString("settings.novelty_report.email"),
				DataBase:    viper.GetString("settings.novelty_report.database"),
			},
		}
	})
}

type NoveltyReportConfig struct {
	DepartName  string
	ContactAddr string
	ZipCode     string
	ManagerName string
	ManagerTel  string
	ContactName string
	ContactTel  string
	Email       string
	DataBase    string
}
