package service

import (
	"github.com/gin-gonic/gin"

	"login-api/db"
	"login-api/entity"
)

type Service struct{}

type User entity.User
type Memo entity.Memo

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

// GetByname is get
func (s Service) GetByName(username string) (User, error) {
	db := db.GetDB()
	var b User

	if err := db.Where("name = ?", username).First(&b).Error; err != nil {
		return b, err
	}

	return b, nil
}

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

func (s Service) MemoGetAll() ([]Memo, error) {
	db := db.GetDB()
	var m []Memo

	if err := db.Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (s Service) MemoCreateModel(c *gin.Context) (Memo, error) {
	db := db.GetDB()
	var m Memo

	if err := c.BindJSON(&m); err != nil {
		return m, err
	}

	if err := db.Create(&m).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (s Service) MemoDeleteByID(id string) error {
	db := db.GetDB()
	var m Memo

	if err := db.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}
