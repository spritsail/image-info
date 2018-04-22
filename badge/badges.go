package badge

import (
	"github.com/gin-gonic/gin"
	"github.com/spritsail/go-badge"
	mb "github.com/spritsail/image-info/microbadger"
	"github.com/urfave/cli"
	"net/http"
	"strings"
	"time"
)

func BuildRoutes(router *gin.RouterGroup, c *cli.Context) (err error) {
	router.GET("/lastbuild/*repo", lastBuildBadge)
	router.GET("/version/*repo", versionBadge)
	router.GET("/pulls/*repo", pullsBadge)
	router.GET("/stars/*repo", starsBadge)
	return
}

type badgeGen func(*mb.Image, string, string, string) (string, string, string)

func repoInfo(req *gin.Context, color string, text string, handler badgeGen) {
	repo := strings.TrimSuffix(req.Param("repo"), ".svg")

	color = req.DefaultQuery("color", color)
	left := req.DefaultQuery("text", text)
	right := ""

	info, status, err := mb.GetImage(repo)
	if err != nil {
		sendBadge(req, left, "error", "red")
		return
	}

	switch status {
	case http.StatusOK:
		color, left, right = handler(&info, repo, color, left)
	case http.StatusNotFound:
		color = "red"
		right = "not found"
	default:
		color = "red"
		right = "error"
	}

	sendBadge(req, left, right, color)
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
