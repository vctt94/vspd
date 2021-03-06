package webapi

import (
	"time"

	"github.com/gin-gonic/gin"
)

// vspInfo is the handler for "GET /vspinfo".
func vspInfo(c *gin.Context) {
	sendJSONResponse(vspInfoResponse{
		Timestamp:     time.Now().Unix(),
		PubKey:        signPubKey,
		FeePercentage: cfg.VSPFee,
		Network:       cfg.NetParams.Name,
		VspClosed:     cfg.VspClosed,
	}, c)
}
