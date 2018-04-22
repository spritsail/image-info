package badge

import (
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/spritsail/go-badge"
	"github.com/spritsail/image-info/microbadger"
	"net/http"
	"strings"
	"time"
)

func lastBuildBadge(req *gin.Context) {
	repo := strings.TrimSuffix(req.Param("repo"), ".svg")
	info, status, err := api.GetImage(repo)
	if err != nil {
		req.JSON(http.StatusInternalServerError, err)
		return
	}

	color := req.DefaultQuery("color", "blue")
	left := req.DefaultQuery("text", "last build")
	right := "error"

	switch status {
	case http.StatusOK:
		parts := strings.Split(repo, ":")
		if len(parts) > 1 {
			ver := api.GetTag(parts[1], &info)
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
