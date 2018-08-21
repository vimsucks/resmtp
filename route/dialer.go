package route

import (
	"github.com/gin-gonic/gin"
	"github.com/vimsucks/resmtp/model"
	"net/http"
	"github.com/storozhukBM/verifier"
	"github.com/vimsucks/resmtp/util"
)


func CreateDialer(ctx *gin.Context) {
	var dl model.Dialer
	if err := ctx.BindJSON(&dl); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verify := verifier.New()
	verify.That(dl.Host != "", "dialer Host cannot be empty")
	verify.That(dl.Port != 0, "dialer Port must be set")
	verify.That(dl.Username != "", "dialer Username cannot be empty")
	verify.That(dl.Password != "", "dialer Password cannot be empty")
	if err := verify.GetError(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := util.GetUniqueDialerAccessToken(&dl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dl.AccessToken = accessToken
	err = model.CreateDialer(&dl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": dl})
}