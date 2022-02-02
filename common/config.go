package common

type Config struct {
	GinMode  string `json:"GIN_MODE"`
	Database Database
	Port     string `json:"PORT"`
}
