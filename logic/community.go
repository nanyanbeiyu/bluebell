package logic

import (
	"bluebell/api"
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

type CommunityLg struct {
}

var communityDao *mysql.CommunityDao

// GetCommunityList 获取社区列表
func (c *CommunityLg) GetCommunityList() ([]*models.GetCommunityListParams, error) {
	// 查询数据库
	communityList, err := communityDao.GetCommunityList()
	if err != nil {
		return nil, err
	}
	return communityList, nil
}

// GetCommunityDetail 获取社区详情
func (c *CommunityLg) GetCommunityDetail(id int64) (*models.Community, int) {
	// 1. 查询社区是否存在
	exitCommunity := communityDao.ExitCommunity("", id)
	if exitCommunity {
		return nil, api.CommunityNotExist
	}
	community, code := communityDao.GetCommunityDetail(id)
	if code != api.CodeSuccess {
		return nil, code
	}
	return community, api.CodeSuccess
}

// CaretCommunity 创建社区
func (c *CommunityLg) CaretCommunity(data *models.CaretCommunityParams) (code int) {
	// 1. 查询社区是否存在
	exitCommunity := communityDao.ExitCommunity(data.Community_name, 0)
	if !exitCommunity {
		return api.CommunityExist
	}

	//2. 生成社区ID
	communityID := snowflake.GenID()

	//3. 插入数据库
	err := communityDao.InsertCommunity(data, communityID)
	if err != nil {
		return api.CodeServerBusy
	}
	return api.CodeSuccess
}

// DeleteCommunityByID 删除社区
func (c *CommunityLg) DeleteCommunityByID(id int64) (code int) {
	// 1. 查询社区是否存在
	exitCommunity := communityDao.ExitCommunity("", id)
	if exitCommunity {
		return api.CommunityNotExist
	}
	// 2. 删除社区
	flag := communityDao.DeleteCommunityByID(id)
	if !flag {
		return api.CodeServerBusy
	}
	return api.CodeSuccess
}
