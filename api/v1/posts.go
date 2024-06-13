package v1

import (
	"bluebell/api"
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var PostLg *logic.PostLg

// CaretPostHandler 创建帖子
// @Summary 创建帖子
// @Description 创建帖子
// @Tags posts
// @Accept application/json
// @Produce application/json
// @Param id body models.CartePostParams true "创建帖子"
// @Security ApiKeyAuth
// @Router /posts [post]
func CaretPostHandler(c *gin.Context) {
	// 1. 获取参数并校验
	data := new(models.CartePostParams)
	if err := c.ShouldBindJSON(data); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("CaretPost with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			api.ResponseError(c, api.CodeInvalidParam)
			return
		}
		api.ResponseErrorWithMsg(c, api.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 业务处理
	// 2.1 获取当前用户
	id, err := api.GetUserID(c)
	if err != nil {
		api.ResponseError(c, api.CodeNeedLogin)
		return
	}
	code := PostLg.CreatePost(data, id)
	if code != api.CodeSuccess {
		api.ResponseError(c, code)
		return
	}
	api.ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情
// @Summary 获取帖子详情
// @Description 获取帖子详情
// @Tags posts
// @Accept application/json
// @Produce application/json
// @Param id path int true "帖子id"
// @Security ApiKeyAuth
// @Router /posts/{id} [get]
func GetPostDetailHandler(c *gin.Context) {
	// 1.获取参数（从URL中获取帖子id）
	postID := c.Param("id")
	pid, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		zap.L().Error("get post id err", zap.Error(err))
		api.ResponseError(c, api.CodeInvalidParam)
		return
	}
	// 2.根据id取出帖子数据
	data, code := PostLg.GetPostDetail(pid)
	if code != api.CodeSuccess {
		api.ResponseError(c, code)
		return
	}
	// 3.返回响应
	api.ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表
// @Summary 获取帖子列表
// @Description 获取帖子列表
// @Tags posts
// @Accept application/json
// @Produce application/json
// @Param page query int false "页码"
// @Param size query int false "每页显示的条数"
// @Security ApiKeyAuth
// @Router /posts [get]
func GetPostListHandler(c *gin.Context) {
	var (
		page, size int64
		err        error
	)
	// 1. 获取参数
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	// 2.获取数据
	data, code := PostLg.GetPostList(page, size)
	if code != api.CodeSuccess {
		api.ResponseError(c, code)
		return
	}
	api.ResponseSuccess(c, data)
}

// GetPostListHandler2 升级获取帖子列表
// 根据前端传来的参数动态的获取帖子列表
// 按照 创建时间 / 分数 进行排序
// @Summary 获取帖子列表
// @Description 获取帖子列表
// @Tags posts
// @Accept application/json
// @Produce application/json
// @Param page query int false "页码"
// @Param size query int false "每页显示的条数"
// @Security ApiKeyAuth
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	p := &models.PostListParams{
		Page:  0,
		Size:  0,
		Order: models.OrderTime,
	}

	// 1. 获取参数
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		api.ResponseError(c, api.CodeInvalidParam)
		return
	}
	// 2.获取数据
	//data, code := PostLg.GetPostList2(p)
	data, code := PostLg.GetPostListNew(p) // 更新合二为一
	if code != api.CodeSuccess {
		api.ResponseError(c, code)
		return
	}
	api.ResponseSuccess(c, data)

}

// GetCommunitPostListHandler 获取社区帖子列表
//func GetCommunitPostListHandler(c *gin.Context) {
//	p := &models.PostListParams{
//		PostListParams: models.PostListParams{
//			Page:  0,
//			Size:  0,
//			Order: models.OrderTime,
//		},
//		CommunityID: 0,
//	}
//	// 1. 获取参数
//	if err := c.ShouldBindQuery(p); err != nil {
//		zap.L().Error("GetCommunitPostListHandler with invalid params", zap.Error(err))
//		api.ResponseError(c, api.CodeInvalidParam)
//		return
//	}
//	// 2.获取数据
//	data, code := PostLg.GetCommunityPostList(p)
//	if code != api.CodeSuccess {
//		api.ResponseError(c, code)
//		return
//	}
//	api.ResponseSuccess(c, data)
//}
