package badge

import (
	"github.com/gin-gonic/gin"
	mb "github.com/spritsail/image-info/microbadger"
	"strings"
)

func versionBadge(req *gin.Context) {
	repoInfo(req, "blue", "docker pulls", func(info *mb.Image, repo string, color string, left string) (string, string, string) {
		var labels map[string]string

		// If a tag is specified, use the tag labels instead
		parts := strings.Split(repo, ":")
		if len(parts) > 1 {
			ver := info.FindTag(parts[1])
			labels = ver.Labels
		} else {
			labels = info.Labels
		}

		verLabel := req.DefaultQuery("label", "org.label-schema.version")

		// Take the label if it exists
		if val, ok := labels[verLabel]; ok {
			return color, left, val
		}

		return color, left, "none"
	})
}
