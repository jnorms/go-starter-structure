package routes

import (
	"github.com/gin-gonic/gin"
	"jnorms.dev/common"
)

type Engine struct {
	Engine *gin.Engine
	Config *common.Config
}

func Routes(config *common.Config) *gin.Engine {
	gin.SetMode(config.GinMode)
	routes := gin.Default()

	return routes
}
