package my_config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	CurrentOtherConfig *OtherConfig
	onceConfig         sync.Once
)

type OtherConfig struct {
	FileServerHost string
}

func LoadPatentConfig() {
	onceConfig.Do(func() {
		CurrentOtherConfig = &OtherConfig{
			FileServerHost: viper.GetString("settings.files.url"),
		}
	})
}
