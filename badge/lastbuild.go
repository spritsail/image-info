package badge

import (
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	mb "github.com/spritsail/image-info/microbadger"
	"net/http"
	"strings"
)

func lastBuildBadge(req *gin.Context) {
	repo := strings.TrimSuffix(req.Param("repo"), ".svg")
	info, status, err := mb.GetImage(repo)
	if err != nil {
		req.JSON(http.StatusInternalServerError, err)
		return
	}

	color := req.DefaultQuery("color", "blue")
	left := req.DefaultQuery("text", "last build")
	right := "error"

	switch status {
	case http.StatusOK:
		// If a tag is specified, use the tag Created time instead
		parts := strings.Split(repo, ":")
		if len(parts) > 1 {
			ver := mb.FindTag(parts[1], &info)
			right = humanize.Time(ver.Created)
		} else {
			right = humanize.Time(info.LastUpdated)
		}
	case http.StatusNotFound:
		color = "red"
		right = "not found"
	default:
		color = "red"
	}

	sendBadge(req, left, right, color)
}
