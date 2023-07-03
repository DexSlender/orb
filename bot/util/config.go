package util

import (
	"strings"

	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
	"github.com/spf13/viper"
)

func SetupConfig(l log.Logger) *viper.Viper {
	v := viper.NewWithOptions(viper.KeyDelimiter("~"))
	v.SetConfigType("toml")
	v.SetConfigName("botconfig")
	v.AddConfigPath(".")
	
	v.SetDefault("debug-log-level", false)
	v.SetDefault("bot~token", "")
	v.SetDefault("bot~setup-commands", true)
	v.SetDefault("bot~global-commands", false)
	v.SetDefault("bot~guild-ids", []snowflake.ID{})

	// v.SetEnvPrefix("ORB")
	// v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			l.Info("Config file not found, writing in current directory")
			if err := v.SafeWriteConfigAs("botconfig.toml"); err != nil {
				l.Fatal("Error while writing config file: ", err)
			}
		} else {
			l.Fatal("Error while reading config file: ", err)
		}
	}

	// c := Config{Viper: v}
	// if err := v.Unmarshal(&c, func(dc *mapstructure.DecoderConfig) {
	// 	dc.TagName = "cfg"
	// }); err != nil {
	// 	l.Error("Config unmarshaler error: ", err)
	// }

	return v
}
