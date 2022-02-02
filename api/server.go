package api

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"jnorms.dev/common"
	"jnorms.dev/routes"
	"log"
)

type Server struct {
	Config *common.Config
}

func (server *Server) Start() {

	server.loadEnvConfigs()

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
	err = json.Unmarshal(envJson, &server.Config.Database)
}

func (server *Server) loadRoutes() {
	r := &routes.Engine{
		Engine: routes.Routes(server.Config),
		Config: server.Config,
	}

	r.Api()

	r.Engine.Run(":" + r.Config.Port)
}
