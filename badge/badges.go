package badge

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

func BuildRoutes(router *gin.RouterGroup, c *cli.Context) (err error) {
	router.GET("/lastbuild/*repo", lastBuildBadge)
	return
}
