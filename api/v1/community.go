package v1

// --- 社区相关接口 ---

import (
	"bluebell/api"
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var communityLg *logic.CommunityLg

// GetCommunityHandler 获取社区列表
// @Summary 获取社区列表
// @Description 获取社区列表
// @Tags 社区列表
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200 {object} models.GetCommunityListParams "成功"
// @Router /community [get]
// @Request Body models.GetCommunityListParams "社区列表"
func GetCommunityHandler(c *gin.Context) {
	// 获取社区列表 (Community_id, Community_name) list
	list, err := communityLg.GetCommunityList()
	if err != nil {
		api.ResponseErrorWithMsg(c, 200, "获取社区列表失败")
		return
	}
	api.ResponseSuccess(c, list)
}

// GetCommunityDetailHandler 获取社区详情
// @Summary 获取社区详情
// @Description 获取社区详情
// @Tags 社区详情
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200 {object} models.Community "成功"
// @Router /community/:id [get]
func GetCommunityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		api.ResponseError(c, api.CodeInvalidParam)
		return
	}
	community, code := communityLg.GetCommunityDetail(id)
	if code != api.CodeSuccess {
		api.ResponseError(c, code)
		return
	}

	api.ResponseSuccess(c, community)
}

// CaretCommunityHandler 添加社区
// @Summary 添加社区
// @Description 添加社区
// @Tags 社区详情
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200
// @Router /community [post]
func CaretCommunityHandler(c *gin.Context) {
	data := new(models.CaretCommunityParams)
	// 1. 获取请求参数及参数校验
	if err := c.ShouldBindJSON(data); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("CaretCommunity with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			api.ResponseError(c, api.CodeInvalidParam)
			return
		}
		api.ResponseErrorWithMsg(c, api.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 处理业务逻辑
	code := communityLg.CaretCommunity(data)

	if code != api.CodeSuccess {
		api.ResponseError(c, code)
		return
	}
	api.ResponseSuccess(c, nil)
}

// DeleteCommunityHandler 删除社区
// @Summary 删除社区
// @Description 删除社区
// @Tags 社区详情
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200
// @Router /community/:id [delete]
func DeleteCommunityHandler(c *gin.Context) {
	// 1. 获取请求参数
	idStr := c.Param("id")
	// 2. 逻辑操作
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		api.ResponseError(c, api.CodeInvalidParam)
		return
	}
	code := communityLg.DeleteCommunityByID(id)
	if code != api.CodeSuccess {
		api.ResponseError(c, code)
		return
	}
	api.ResponseSuccess(c, nil)
}
