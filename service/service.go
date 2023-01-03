package service

import (
	"github.com/gin-gonic/gin"

	"login-api/db"
	"login-api/entity"
)

type Service struct{}

type User entity.User

// GetAll is get all
func (s Service) GetAll() ([]User, error) {
	db := db.GetDB()
	var b []User

	if err := db.Find(&b).Error; err != nil {
		return nil, err
	}

	return b, nil
}

// CreateModel is create
func (s Service) CreateModel(c *gin.Context) (User, error) {
	db := db.GetDB()
	var b User

	if err := c.BindJSON(&b); err != nil {
		return b, err
	}

	if err := db.Create(&b).Error; err != nil {
		return b, err
	}

	return b, nil
}

// GetByID is get
func (s Service) GetByID(id string) (User, error) {
	db := db.GetDB()
	var b User

	if err := db.Where("id = ?", id).First(&b).Error; err != nil {
		return b, err
	}

	return b, nil
}

// // GetByDantaimei is get
// func (s Service) GetByDantaimei(dantai string) ([]User, error) {
// 	db := db.GetDB()
// 	var b []User

// 	if err := db.Where("dantai LIKE ?", "%"+dantai+"%").Find(&b).Error; err != nil {
// 		return b, err
// 	}

// 	return b, nil
// }

// UpdateByID is update
func (s Service) UpdateByID(id string, c *gin.Context) (User, error) {
	db := db.GetDB()
	var b User

	if err := db.Where("id = ?", id).First(&b).Error; err != nil {
		return b, err
	}

	if err := c.BindJSON(&b); err != nil {
		return b, err
	}

	db.Save(&b)

	return b, nil
}

// DeleteByID is delete
func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var b User

	if err := db.Where("id = ?", id).Delete(&b).Error; err != nil {
		return err
	}

	return nil
}
