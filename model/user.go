package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Nickname  string     `json:"nickname"`
	Mobile    string     `json:"mobile"`
	Gender    int32      `json:"gender"` // 性别：0，未知；1，男；2，女；
	Birthday  *time.Time `json:"birthday"`
	IsDel     int32      `json:"is_del"` // 是否删除：0，默认；1，已删除；2，未删除；
	DeletedAt *time.Time `json:"deleted_at"`
	CreatedAt time.Time  `json:"created_at"` // 创建时间
	UpdatedAt time.Time  `json:"updated_at"` // 更新时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

func GetUserByID(db *gorm.DB, id int64) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func GetUserList(db *gorm.DB, page int32, pageSize int32) (list []*User, count int64) {
	db.Model(&User{}).Count(&count).Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Order("id desc").Find(&list)
	return
}
