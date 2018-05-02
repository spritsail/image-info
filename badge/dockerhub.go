package badge

import (
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	mb "github.com/spritsail/image-info/microbadger"
)

func pullsBadge(req *gin.Context) {
	repoInfo(req, "blue", "docker pulls", func(info *mb.Image, repo string, color string, left string) (string, string, string) {
		return color, left, humanize.SIWithDigits(float64(info.PullCount), 1, "")
	})
}
func starsBadge(req *gin.Context) {
	repoInfo(req, "blue", "docker stars", func(info *mb.Image, repo string, color string, left string) (string, string, string) {
		return color, left, humanize.SIWithDigits(float64(info.StarCount), 1, "")
	})
}
