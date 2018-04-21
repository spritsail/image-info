package badge

import (
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/spritsail/go-badge"
	"github.com/spritsail/image-info/microbadger"
	"net/http"
)

func lastBuildBadge(req *gin.Context) {
	owner := req.Param("owner")
	imgtag := req.Param("image")

	info, status, err := api.GetImage(owner + "/" + imgtag)
	if err != nil {
		req.JSON(http.StatusInternalServerError, err)
		return
	}

	color := req.DefaultQuery("color", "blue")
	left := req.DefaultQuery("text", "last build")
	right := "error"

	switch status {
	case http.StatusOK:
		right = humanize.Time(info.LastUpdated)
	case http.StatusNotFound:
		color = "red"
		right = "not found"
	default:
		color = "red"
		// req.JSON(http.StatusOK, info)
	}

	badge, err := badge.RenderDef(left, right, badge.Color(color))
	if err != nil {
		req.AbortWithStatus(500)
		return
	}
	req.Header("Content-Type", "image/svg+xml")
	req.String(http.StatusOK, badge)
}
