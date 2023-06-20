package models

// 定义 user 的 dao 层操作
import (
	"gorm.io/gorm"
)

type User struct {
	UserId   uint   `gorm:"column:user_id"`
	UserName string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	Salt     string `gorm:"column:salt"`
}

// 根据 username 查找单个用户
func GetUser(db *gorm.DB, username string) (user *User, exists bool, err error) {
	err = db.Where("user_name=?", username).Limit(1).Find(&user).Error
	if err != nil {
		// 查询发生错误
		return nil, false, err
	}

	// 未查询到数据
	if user.UserId == 0 {
		return nil, false, nil
	}

	return user, true, err
}

// 添加用户
func AddUser(db *gorm.DB, user *User) (err error) {
	if err = db.Select("UserName", "Password", "Salt").Create(user).Error; err != nil {
		return err
	}
	return nil
}

// 根据用户名修改密码
func UpdatePasswd(db *gorm.DB, user *User) (err error) {
	if err = db.Model(user).Select("Password", "Salt").Where("user_name=?", user.UserName).
		Updates(User{Password: user.Password, Salt: user.Salt}).Error; err != nil {
		return err
	}
	return nil
}
