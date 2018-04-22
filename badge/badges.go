package badge

import (
	"github.com/gin-gonic/gin"
	"github.com/spritsail/go-badge"
	"github.com/urfave/cli"
	"net/http"
	"time"
)

func BuildRoutes(router *gin.RouterGroup, c *cli.Context) (err error) {
	router.GET("/lastbuild/*repo", lastBuildBadge)
	return
}

func sendBadge(req *gin.Context, left string, right string, color string) {
	badge, err := badge.RenderDef(left, right, badge.Color(color))
	if err != nil {
		req.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	req.Header("Content-Type", "image/svg+xml")
	req.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	req.Header("Expires", time.Now().UTC().Format(time.RFC1123Z))
	req.String(http.StatusOK, badge)
}
