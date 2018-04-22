package badge

import (
	"github.com/gin-gonic/gin"
	mb "github.com/spritsail/image-info/microbadger"
	"net/http"
	"strings"
)

func versionBadge(req *gin.Context) {
	repo := strings.TrimSuffix(req.Param("repo"), ".svg")
	info, status, err := mb.GetImage(repo)
	if err != nil {
		req.JSON(http.StatusInternalServerError, err)
		return
	}

	verLabel := req.DefaultQuery("label", "org.label-schema.version")
	color := req.DefaultQuery("color", "blue")
	left := req.DefaultQuery("text", "version")
	right := "none"

	switch status {
	case http.StatusOK:
		var labels map[string]string

		// If a tag is specified, use the tag labels instead
		parts := strings.Split(repo, ":")
		if len(parts) > 1 {
			ver := info.FindTag(parts[1])
			labels = ver.Labels
		} else {
			labels = info.Labels
		}

		// Take the label if it exists
		if val, ok := labels[verLabel]; ok {
			right = val
		}
	case http.StatusNotFound:
		color = "red"
		right = "not found"
	default:
		color = "red"
		right = "error"
	}

	sendBadge(req, left, right, color)
}
