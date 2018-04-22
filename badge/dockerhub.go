package badge

import (
	"github.com/gin-gonic/gin"
	mb "github.com/spritsail/image-info/microbadger"
	"github.com/spritsail/numtext"
)

func pullsBadge(req *gin.Context) {
	repoInfo(req, "blue", "docker pulls", func(info *mb.Image, repo string, color string, left string) (string, string, string) {
		return color, left, numtext.ToText(info.PullCount)
	})
}
func starsBadge(req *gin.Context) {
	repoInfo(req, "blue", "docker stars", func(info *mb.Image, repo string, color string, left string) (string, string, string) {
		return color, left, numtext.ToText(info.StarCount)
	})
}
