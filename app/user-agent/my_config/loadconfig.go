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
	InnojoyUser     string
	InnojoyPassword string
}

func LoadPatentConfig() {
	onceConfig.Do(func() {
		CurrentPatentConfig = &PatentConfig{
			InnojoyUser:     viper.GetString("settings.user-agent.innojoy.user"),
			InnojoyPassword: viper.GetString("settings.user-agent.innojoy.password"),
		}
	})
}
