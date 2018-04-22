package utilities

import "github.com/spf13/viper"

func IsDevelopment() bool {
	if len(viper.GetString("ENV")) == 0 || viper.GetString("ENV") == "development" {
		return true
	}
	return false
}
