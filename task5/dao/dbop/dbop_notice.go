package dbop

import (
	"errors"
	"gorm.io/gorm"
	"mission/task5/dao/model"
)

// 发布公告
func CreateNotice(db *gorm.DB, notice *model.Notice) error {
	result := db.Create(notice)
	return result.Error
}

// 根据id展示指定公告
func ShowNotice(db *gorm.DB, id uint64) ([]*model.Notice, error) {
	var notices []*model.Notice
	result := db.Find(&notices, "id = ?", id)
	return notices, result.Error
}

// 展示所有公告
func ShowAllNotice(db *gorm.DB) ([]*model.Notice, error) {
	var notices []*model.Notice
	result := db.Find(&notices)
	return notices, result.Error
}

// 更新公告
func UpdateNotice(db *gorm.DB, notice *model.Notice, id uint64) error {
	result := db.Model(&model.Notice{}).Where("id = ?", id).Updates(notice)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("更新公告失败,未找到公告")
	}
	return result.Error
}

// 删除公告
func DeleteNotice(db *gorm.DB, id uint64) error {
	result := db.Where("id = ?", id).Delete(&model.Notice{})
	return result.Error
}
