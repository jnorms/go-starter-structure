package common

type Config struct {
	GinMode  string `json:"GIN_MODE"`
	Database Database
	Port     string `json:"APP_PORT"`
	Timezone string `json:"APP_TIMEZONE"`
}
