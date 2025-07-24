package config

import "os"

var Config = struct {
	ApiKey         string
	Port           string
	RaribleRootURL string
}{
	ApiKey:         os.Getenv("API_KEY"),
	Port:           os.Getenv("PORT"),
	RaribleRootURL: os.Getenv("RARIBLE_ROOT_URL"),
}
