package userdao

import (
	"boframe/models/usermodels"
	"boframe/settings/mysqlI"
)

func CheckUsernameExists(username string) (bool, error) {
	var num int64
	db := mysqlI.DB()
	res := db.Model((*usermodels.User)(nil)).Select("id").Where("username=? AND is_delete = false", username).
		Count(&num)
	if res.Error != nil {
		return false, res.Error
	}
	return num > 0, nil
}

func Insert(user *usermodels.User) error {
	db := mysqlI.DB()
	return db.Create(user).Error
}

func OneByUsernameNotDeleted(username string) (user *usermodels.User, err error) {
	user = new(usermodels.User)
	err = mysqlI.DB().Where("username=? AND is_delete = false", username).First(user).Error
	return
}
