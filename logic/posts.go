package logic

import (
	"bluebell/api"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

var postDao *mysql.PostDao

type PostLg struct {
}

// CreatePost 创建帖子
func (p *PostLg) CreatePost(data *models.CartePostParams, uid int64) (code int) {
	// 1.生成帖子ID
	postID := snowflake.GenID()
	// 2.添加到数据库
	if err := postDao.InsertPost(data, uid, postID); err != nil {
		fmt.Println("InsertPost failed, err:", err)
		return api.CodeServerBusy
	}
	err := redis.CreatePost(postID, data.CommunityID)
	if err != nil {
		fmt.Println("CreatePost failed, err:", err)
		return api.CodeServerBusy
	}
	return api.CodeSuccess
}

// GetPostDetail 获取帖子详情
func (p *PostLg) GetPostDetail(id int64) (data *models.ApiPostDetail, code int) {
	detail, err := postDao.GetPostDetail(id)
	if err != nil {
		zap.L().Error("GetPostDetail failed", zap.Error(err))
		return nil, api.CodeServerBusy
	}
	user, err := userDao.GetUserByID(detail.AuthorID)
	if err != nil {
		zap.L().Error("GetUserByID failed", zap.Error(err))
		return nil, api.CodeServerBusy
	}
	communityDetail, _ := communityDao.GetCommunityDetail(detail.CommunityID)
	data = &models.ApiPostDetail{
		Post:           detail,
		AuthorName:     user.User_name,
		Community_name: communityDetail.Community_name,
	}
	return data, api.CodeSuccess
}

// GetPostList 获取帖子列表
func (p *PostLg) GetPostList(page, size int64) (data []*models.ApiPostDetail, code int) {
	posts, err := postDao.GetPostList(page, size)
	if err != nil {
		zap.L().Error("GetPostList failed", zap.Error(err))
		return nil, api.CodeServerBusy
	}
	for _, post := range posts {
		user, err := userDao.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("GetUserByID failed", zap.Error(err))
			continue
		}
		communityDetail, _ := communityDao.GetCommunityDetail(post.CommunityID)
		data = append(data, &models.ApiPostDetail{
			Post:           post,
			AuthorName:     user.User_name,
			Community_name: communityDetail.Community_name,
		})
	}
	return data, api.CodeSuccess
}

// GetPostList2 获取帖子列表2
func (p *PostLg) GetPostList2(d *models.PostListParams) (data []*models.ApiPostDetail, code int) {
	// 2. 去redis查询id列表
	idList, err := redis.GetPostIDByOrder(d)
	if err != nil {
		zap.L().Error("GetPostList2 with GetPostIDByOrder:", zap.Error(err))
		return nil, api.CodeServerBusy
	}
	if len(idList) == 0 {
		zap.L().Warn("redis.GetPostIDByOrder return 0 data")
		return nil, api.CodeServerBusy
	}
	// 3. 根据id列表去数据库中查询相关数据
	posts, err := postDao.GetPostsListByIDs(idList)
	if err != nil {
		zap.L().Error("GetPostList2 with GetPostsListByIDs:", zap.Error(err))
		return nil, api.CodeServerBusy
	}
	// 4. 将帖子的数据与作者信息关联起来
	for _, post := range posts {
		user, err := userDao.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("GetUserByID failed", zap.Error(err))
			continue
		}
		communityDetail, _ := communityDao.GetCommunityDetail(post.CommunityID)
		data = append(data, &models.ApiPostDetail{
			Post:           post,
			AuthorName:     user.User_name,
			Community_name: communityDetail.Community_name,
		})
	}
	return data, api.CodeSuccess
}

// GetCommunityPostList 获取社区帖子列表
func (p *PostLg) GetCommunityPostList(d *models.PostListParams) (data []*models.ApiPostDetail, code int) {
	idList, err := redis.GetCommunityPostIDsInOrder(d)
	if err != nil {
		zap.L().Error("GetCommunityPostIDsInOrder failed", zap.Error(err))
		return nil, api.CodeServerBusy
	}
	if len(idList) == 0 {
		zap.L().Warn("redis.GetPostIDByOrder return 0 data")
		return nil, api.CodeServerBusy
	}
	// 3. 根据id列表去数据库中查询相关数据
	posts, err := postDao.GetPostsListByIDs(idList)
	if err != nil {
		zap.L().Error("GetPostList2 with GetPostsListByIDs:", zap.Error(err))
		return nil, api.CodeServerBusy
	}
	// 4. 将帖子的数据与作者信息关联起来
	for _, post := range posts {
		user, err := userDao.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("GetUserByID failed", zap.Error(err))
			continue
		}
		communityDetail, _ := communityDao.GetCommunityDetail(post.CommunityID)
		data = append(data, &models.ApiPostDetail{
			Post:           post,
			AuthorName:     user.User_name,
			Community_name: communityDetail.Community_name,
		})
	}
	return data, api.CodeSuccess
}

// GetPostListNew 将两个查询逻辑合二为一的函数
func (p *PostLg) GetPostListNew(d *models.PostListParams) (data []*models.ApiPostDetail, code int) {
	// 根据请求参数的不同，执行不同的逻辑
	if d.CommunityID == 0 {
		// 查询所有
		data, code = p.GetPostList2(d)
	} else {
		// 根据社区id查询
		data, code = p.GetCommunityPostList(d)
	}
	return
}
