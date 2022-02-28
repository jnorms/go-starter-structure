package common

const (
	DB_DRIVER_MYSQL    = "mysql"
	DB_DRIVER_POTSGRES = "postgres"
)

type Database struct {
	Driver   string `json:"DB_DRIVER"`
	Host     string `json:"DB_HOST"`
	Name     string `json:"DB_NAME"`
	User     string `json:"DB_USER"`
	Password string `json:"DB_PASSWORD"`
	Port     string `json:"DB_PORT"`
	Ssl      string `json:"DB_SSL"`
	Timezone string `json:"DB_TIMEZONE"`
}
