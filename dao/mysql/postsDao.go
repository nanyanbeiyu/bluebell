package mysql

import (
	"bluebell/models"
	"fmt"
	"strings"
)

type PostDao struct {
}

// InsertPost 插入帖子
func (p *PostDao) InsertPost(data *models.CartePostParams, uid int64, postID int64) (err error) {
	post := &models.Post{
		Title:       data.Title,
		Content:     data.Content,
		AuthorID:    uid,
		CommunityID: data.CommunityID,
		PostID:      postID,
	}
	if err = db.Create(post).Error; err != nil {
		return err
	}
	return
}

// GetPostDetail 获取帖子详情
func (p *PostDao) GetPostDetail(id int64) (data *models.Post, err error) {
	data = new(models.Post)
	err = db.Where("post_id = ?", id).First(data).Error
	if err != nil {
		return nil, err
	}
	return
}

// GetPostList 获取帖子列表
func (p *PostDao) GetPostList(page, size int64) (data []*models.Post, err error) {
	postList := make([]*models.Post, 0, size)
	err = db.Order("created_at desc").Offset(int((page - 1) * size)).Limit(int(size)).Find(&postList).Error
	if err != nil {
		return nil, err
	}
	return postList, nil
}

func (p *PostDao) GetPostsListByIDs(ids []string) ([]*models.Post, error) {
	var posts []*models.Post
	sql1 := fmt.Sprintf("select * from post where post_id in (%s) ORDER BY FIELD(post_id,%s)", strings.Join(ids, ","), strings.Join(ids, ","))
	err := db.Raw(sql1).Scan(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
