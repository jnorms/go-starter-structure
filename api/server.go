package api

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"jnorms.dev/common"
	"jnorms.dev/database"
	"jnorms.dev/routes"
	"log"
)

type Server struct {
	Config *common.Config
	DB     *gorm.DB
}

func (server *Server) Start() {
	server.loadEnvConfigs()

	server.startDBConnection()

	server.loadRoutes()
}

func (server *Server) loadEnvConfigs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env, err := godotenv.Read()

	if err != nil {
		log.Fatal("Error reading .env file")
	}

	envJson, err := json.Marshal(env)

	if err != nil {
		log.Fatal("Error marshal .env file")
	}

	err = json.Unmarshal(envJson, &server.Config)

	if err != nil {
		log.Fatal("Error unmarshal env json")
	}

	err = json.Unmarshal(envJson, &server.Config.Database)

	if err != nil {
		log.Fatal("Error unmarshal env json database config", err)
	}
}

func (server *Server) loadRoutes() {
	r := &routes.Engine{
		Engine: routes.Routes(server.Config),
		Config: server.Config,
	}

	r.Api()

	r.Engine.Run(":" + r.Config.Port)
}

func (server *Server) startDBConnection() {
	db := &database.Database{
		Database: &server.Config.Database,
	}

	server.DB = db.Connect()
}
