package route

import (
	"github.com/vimsucks/resmtp/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/vimsucks/resmtp/mail"
	"github.com/storozhukBM/verifier"
)

type MailRequestBody struct {
	Message model.Message `json:"message"`
	AccessToken string `json:"access_token"`
}

func SendMail(ctx *gin.Context) {
	reqBody := MailRequestBody{}
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verify := verifier.New()
	verify.That(len(reqBody.Message.To) != 0, "To cannot be empty")
	verify.That(reqBody.Message.Subject != "", "Subject cannot be empty")
	verify.That(reqBody.Message.Content != "", "Content cannot be empty")
	if err := verify.GetError(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dl, err := model.GetDialerByAccessToken(reqBody.AccessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if dl == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "access token 无效"})
		return
	}

	err = mail.SendMail(&reqBody.Message, dl)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"content": reqBody})
}