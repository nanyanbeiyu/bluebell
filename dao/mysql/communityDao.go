package mysql

import (
	"bluebell/api"
	"bluebell/models"
	"fmt"

	"go.uber.org/zap"
)

type CommunityDao struct {
}

// GetCommunityList 获取社区列表
func (c *CommunityDao) GetCommunityList() (communityList []*models.GetCommunityListParams, err error) {
	err = db.Model(&models.Community{}).Select("community_id,community_name,introduction").Find(&communityList).Error
	if err != nil {
		zap.L().Error("mysql get community list failed", zap.Error(err))
		return nil, err
	}
	return
}

// GetCommunityDetail 获取社区详情
func (c *CommunityDao) GetCommunityDetail(id int64) (community *models.Community, code int) {
	community = new(models.Community)
	err := db.Model(&models.Community{}).Where("community_id = ?", id).Preload("Posts").First(community).Error
	if err != nil {
		zap.L().Error("mysql get community detail failed", zap.Error(err))
		return nil, api.CodeServerBusy
	}
	return community, api.CodeSuccess
}

// InsertCommunity 插入社区
func (c *CommunityDao) InsertCommunity(data *models.CaretCommunityParams, community_id int64) (err error) {
	community := models.Community{
		Community_id:   community_id,
		Community_name: data.Community_name,
		Introduction:   data.Introduction,
	}
	err = db.Create(&community).Error
	if err != nil {
		zap.L().Error("mysql insert community failed", zap.Error(err))
		return err
	}
	return
}

// DeleteCommunityByID 删除社区
func (c *CommunityDao) DeleteCommunityByID(id int64) bool {
	var community models.Community
	err := db.Where("community_id = ?", id).Delete(&community).Error
	if err != nil {
		zap.L().Error("mysql delete community failed", zap.Error(err))
		return false
	}
	return true
}

// ExitCommunity 社区是否存在
func (c *CommunityDao) ExitCommunity(name string, id int64) bool {
	var community models.Community
	var count int64
	if id != 0 {
		db.Where("community_id = ?", id).Find(&community).Count(&count)
	}
	if name != "" {
		db.Where("community_name = ?", name).Find(&community).Count(&count)
	}
	if count == 0 {
		return true
	}
	fmt.Println(count)
	return false
}
