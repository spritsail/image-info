package badge

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	mb "github.com/spritsail/image-info/microbadger"
	"strings"
)

func imageBadge(req *gin.Context) {
	repoInfo(req, "blue", "", func(info *mb.Image, repo string, color string, left string) (string, string, string) {
		// If a tag is specified, use the tag Created time instead
		parts := strings.Split(repo, ":")
		if len(parts) > 1 {
			ver := info.FindTag(parts[1])
			return color, humanize.IBytes(ver.DownloadSize), fmt.Sprintf("%d layers", ver.LayerCount)
		} else {
			return color, humanize.IBytes(info.DownloadSize), fmt.Sprintf("%d layers", info.LayerCount)
		}
	})
}

func commitBadge(req *gin.Context) {
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
