package dbop

import (
	"errors"
	"gorm.io/gorm"
	"mission/task3/model"
)

// 创建用户
func CreateGamer(db *gorm.DB, gamer *model.Gamer) error {
	result := db.Create(gamer)
	return result.Error
}

// 读取用户（根据 Name）
func GetGamerByName(db *gorm.DB, name string) (*model.Gamer, error) {
	var gamer model.Gamer
	result := db.Model(&model.Gamer{}).Where("name = ?", name).First(&gamer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // 如果没有找到记录，返回 nil 而不是错误
		}
		return nil, result.Error
	}
	return &gamer, nil
}

// 查询所有用户
func GetAllGamer(db *gorm.DB) ([]*model.Gamer, error) {
	var gamers []*model.Gamer
	result := db.Find(&gamers)
	return gamers, result.Error
}

// 更新用户（根据 Name）
func UpdateGamer(db *gorm.DB, name string, email string) error {
	result := db.Model(&model.Gamer{}).Where("name=?", name).Update("email", email)
	return result.Error
}

// 删除用户（根据 Name）
func DeleteGamer(db *gorm.DB, name string) error {
	result := db.Where("name = ?", name).Delete(&model.Gamer{})
	return result.Error
}
