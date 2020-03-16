package utilities

import "github.com/spf13/viper"

// IsDevelopment return true when the current setup is development
func IsDevelopment() bool {
	if len(viper.GetString("ENV")) == 0 || viper.GetString("ENV") == "development" {
		return true
	}
	return false
}

// Contains checks if value in array
func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
