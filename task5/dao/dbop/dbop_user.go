package dbop

import (
	"gorm.io/gorm"
	"mission/task5/dao/model"
)

// 创建用户
func CreateUser(db *gorm.DB, user *model.User) error {
	result := db.Create(user)
	return result.Error
}

//// 读取用户（根据 Name）
//func GetUserByName(db *dao.DB, name string) (*model.User, error) {
//	var user model.User
//	result := db.Model(&model.User{}).Where("name = ?", name).First(&user)
//	if result.Error != nil {
//		if errors.Is(result.Error, dao.ErrRecordNotFound) {
//			return nil, nil // 如果没有找到记录，返回 nil 而不是错误
//		}
//		return nil, result.Error
//	}
//	return &user, nil
//}

// 更新用户（根据 Name）
func UpdateUser(db *gorm.DB, name string, email string) error {
	result := db.Model(&model.User{}).Where("name=?", name).Update("email", email)
	return result.Error
}

// 删除用户（根据 Name）
func DeleteUser(db *gorm.DB, name string) error {
	result := db.Where("name = ?", name).Delete(&model.User{})
	return result.Error
}
