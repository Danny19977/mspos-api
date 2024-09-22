package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Fullname        string `gorm:"not null" json:"fullname"`
	Email           string `json:"email" gorm:"unique; not null"`
	Phone           string `json:"phone"`
	Title           string `json:"title"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" gorm:"-"`

	AreaID     uint `gorm:"foreignKey:area_id" json:"area_id"`
	ProvinceID uint `gorm:"foreignKey:province_id" json:"province_id"`
	SupID      uint `gorm:"foreignKey:sup_id" json:"sup_id"`
	// PosID      uint `gorm:"foreignKey:pos_id" json:"pos_id"`

	Role       string `json:"role"`
	Permission string `json:"permission"`
	Image      string `json:"image"`
	Status     bool   `json:"status"`
	IsManager  bool   `json:"is_manager"`
	Signature  string `json:"signature"`

	UserLogs []UserLogs
}

type UserResponse struct {
	Id       uint   `json:"id,omitempty"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Title    string `json:"title"`
	Role     string `json:"role"`
	Area     uint   `json:"area_id"`
	Province uint   `json:"province_id"`
	Sup      uint   `json:"sup_id"`
	// Pos        uint      `json:"pos_id"`
	Permission string    `json:"permission"`
	Status     bool      `json:"status"`
	Signature  string    `json:"signature"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserPaginate struct {
	Id         uint      `json:"id,omitempty"`
	Fullname   string    `json:"fullname"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Title      string    `json:"title"`
	Role       string    `json:"role"`
	Area       string    `json:"area"`
	Province   string    `json:"province"`
	Sup        string    `json:"sup"`
	Permission string    `json:"permission"`
	Status     bool      `json:"status"`
	Signature  string    `json:"signature"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u *User) SetPassword(p string) {
	hp, _ := bcrypt.GenerateFromPassword([]byte(p), 14)
	u.Password = string(hp)
}

func (u *User) ComparePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	return err
}

func (u *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)
	return total
}

func (u *User) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	su := []User{}
	db.Preload("Province").Offset(offset).Limit(limit).Find(&su)
	return su
}
