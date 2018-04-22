package badge

import (
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	mb "github.com/spritsail/image-info/microbadger"
	"strings"
)

func lastBuildBadge(req *gin.Context) {
	repoInfo(req, "blue", "last build", func(info *mb.Image, repo string, color string, left string) (string, string, string) {
		// If a tag is specified, use the tag Created time instead
		parts := strings.Split(repo, ":")
		if len(parts) > 1 {
			ver := info.FindTag(parts[1])
			return color, left, humanize.Time(ver.Created)
		} else {
			return color, left, humanize.Time(info.LastUpdated)
		}
	})
}
