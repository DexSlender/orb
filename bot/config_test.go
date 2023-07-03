package main

import (
	"testing"

	"github.com/disgoorg/snowflake/v2"
	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	viper.SetConfigType("toml")
	viper.SetConfigName("botconfig")
	viper.AddConfigPath(".")

	viper.SetDefault("debug-log-level", false)
	viper.SetDefault("bot~token", "")
	viper.SetDefault("bot~setup-commands", true)
	viper.SetDefault("bot~global-commands", false)
	viper.SetDefault("bot~guild-ids", []snowflake.ID{})

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			t.Log("Config file not found, writing in current directory")
			if err := viper.SafeWriteConfigAs("botconfig.toml"); err != nil {
				t.Fatal("Error while writing config file: ", err)
			}
		} else {
			t.Fatal("Error while reading config file: ", err)
		}
	}
	
	t.Log(viper.GetString("bot~token"))
	/*
		if err := v.Unmarshal(&c, func(dc *mapstructure.DecoderConfig) {
			dc.TagName = "cfg"
			dc.MatchName("bot.token", "token")
		}); err != nil {
			l.Error("Config unmarshaler error: ", err)
		}
	*/
}
