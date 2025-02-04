package service

import (
	"errors"
	"github.com/akazwz/go-gin-restful-api/global"
	"github.com/akazwz/go-gin-restful-api/model"
	"github.com/akazwz/go-gin-restful-api/model/request"
	"github.com/akazwz/go-gin-restful-api/pkg/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Register(u model.User) (err error, userInter *model.User) {
	var user model.User
	if !errors.Is(global.GDB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("username already exist"), userInter
	}
	// md5 password
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GDB.Create(&u).Error
	return err, &u
}

func RegisterByPhoneVerificationCode(u model.User) (err error, userInter *model.User) {
	var user model.User
	global.GDB.Where("phone = ?", u.Phone).First(&user)
	if len(user.Phone) > 1 {
		return errors.New("phone already exist"), userInter
	}
	u.UUID = uuid.NewV4()
	err = global.GDB.Create(&u).Error
	return err, &u
}

func RegisterByOpenId(u model.User) (err error, userInter *model.User) {
	var user model.User
	global.GDB.Where("open_id = ?", u.OpenId).First(&user)
	if len(user.OpenId) > 1 {
		return errors.New("open id already exist"), userInter
	}
	u.UUID = uuid.NewV4()
	err = global.GDB.Create(&u).Error
	return err, &u
}

func IsPhoneExist(phone string) (bool, *model.User) {
	var user model.User
	_ = global.GDB.Where("phone = ?", phone).First(&user).Error
	if len(user.Phone) > 1 {
		return true, &user
	}
	return false, &user
}

func IsOpenIdExist(openId string) (bool, *model.User) {
	var user model.User
	_ = global.GDB.Where("open_id = ?", openId).First(&user).Error
	if len(user.OpenId) > 1 {
		return true, &user
	}
	return false, &user
}

func LoginByUsernamePwd(u *model.User) (err error, userInter *model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

func LoginByPhonePwd(u *model.User) (err error, userInter *model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GDB.Where("phone = ? AND password = ?", u.Phone, u.Password).First(&user).Error
	return err, &user
}

/*func LoginByPhoneCode(phone, code str) (err error, userInter *model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}*/

func ChangePassword(u *model.User, newPassword string) (err error, userInter *model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GDB.Where("uuid = ? AND password = ?", u.UUID, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

func ChangePasswordByPhoneVerificationCode(u *model.User, newPassword string) (err error, userInter *model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GDB.Where("uuid = ?", u.UUID).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

func UpdateUserProfileByUser(u *model.User) (err error, userInter *model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GDB.Where("uuid = ?", u.UUID).First(&user).Updates(model.User{
		NickName:  u.NickName,
		AvatarUrl: u.AvatarUrl,
		Gender:    u.Gender,
		Bio:       u.Bio,
	}).Error
	return err, u
}

func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GDB.Model(&model.User{})
	var userList []model.User
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return err, userList, total
}

func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.GDB.Where("uuid = ?", uuid).First(&model.User{}).Update("authority_id", authorityId).Error
	return err
}

func FindUserByID(id int) (err error, user *model.User) {
	var u model.User
	err = global.GDB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

func FindUserByUUID(uuid string) (err error, user *model.User) {
	var u model.User
	if err = global.GDB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("user does not exist"), &u
	}
	return nil, &u
}
