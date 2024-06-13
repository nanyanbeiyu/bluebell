package v1

import (
	"bluebell/api"
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var voteLg *logic.VoteLg

// PostVoteHandler 投票
// @Summary 投票
// @Description 投票
// @Tags vote
// @Accept application/json
// @Produce application/json
// @Param data body models.VoteData true "投票信息"
// @Security ApiKeyAuth
// @Success 200 {object} api.ResponseData
// @Router /vote [post]
func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(models.VoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("vote with invalid params", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			api.ResponseError(c, api.CodeInvalidParam)
			return
		}
		api.ResponseErrorWithMsg(c, api.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 获取当前用户id
	uid, err := api.GetUserID(c)
	if err != nil {
		api.ResponseError(c, api.CodeNeedLogin)
		return
	}
	err = voteLg.VoteForPost(uid, p)
	if err != nil {
		zap.L().Error("vote failed", zap.Error(err))
		api.ResponseError(c, api.CodeServerBusy)
		return
	}
	api.ResponseSuccess(c, "123")
}
