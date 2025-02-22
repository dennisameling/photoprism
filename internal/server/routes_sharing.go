package server

import (
	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism/internal/api"

	"github.com/photoprism/photoprism/internal/config"
)

// registerSharingRoutes configures link sharing routes.
func registerSharingRoutes(router *gin.Engine, conf *config.Config) {
	s := router.Group(conf.BaseUri("/s"))
	{
		api.Shares(s)
		api.SharePreview(s)
	}
}
