package service

import (
	"github.com/gin-gonic/gin"

	"login-api/db"
	"login-api/entity"
)

type Service struct{}

type User entity.User
type Memo entity.Memo
type Coin entity.Coin

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

func (s Service) MemoDeleteByID(username string) error {
	db := db.GetDB()
	var m Memo

	if err := db.Where("name = ?", username).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s Service) MemoGetByName(username string) ([]Memo, error) {
	db := db.GetDB()
	var m []Memo

	if err := db.Where("name = ?", username).Find(&m).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (s Service) TicketAdd(username string) {
	db := db.GetDB()

	b, _ := s.GetByName(username)
	ti := b.Ticket + 1

	db.Model(b).Update("Ticket", ti)
}

// func (s Service) CoinAdd(username string) {
// 	db := db.GetDB()

// 	b, _ := s.GetByName(username)
// 	co := b.Coin + 1

// 	db.Model(b).Update("Coin", co)
// }

func (s Service) TicketSub(username string) {
	db := db.GetDB()

	b, _ := s.GetByName(username)
	ti := b.Ticket - 1

	db.Model(b).Update("Ticket", ti)
}

// func (s Service) CoinSub(username string) {
// 	db := db.GetDB()

// 	b, _ := s.GetByName(username)
// 	co := b.Coin - 1

// 	db.Model(b).Update("Coin", co)
// }

func (s Service) GameCoinGetAll() ([]Coin, error) {
	db := db.GetDB()
	var co []Coin

	if err := db.Find(&co).Error; err != nil {
		return nil, err
	}

	return co, nil
}

func (s Service) GameCoinCreateModel(c *gin.Context) (Coin, error) {
	db := db.GetDB()
	var co Coin

	if err := c.BindJSON(&co); err != nil {
		return co, err
	}

	if err := db.Create(&co).Error; err != nil {
		return co, err
	}

	return co, nil
}

func (s Service) GameCoinGetByName(username string) (Coin, error) {
	db := db.GetDB()
	var co Coin

	if err := db.Where("name = ?", username).Find(&co).Error; err != nil {
		return co, err
	}

	return co, nil
}

func (s Service) GameCoinUpdateByName(username string, c *gin.Context) (Coin, error) {
	db := db.GetDB()
	var co Coin

	if err := db.Where("name = ?", username).First(&co).Error; err != nil {
		return co, err
	}

	if err := c.BindJSON(&co); err != nil {
		return co, err
	}

	db.Save(&co)

	return co, nil
}

func (s Service) GameCoinDeleteByName(username string) error {
	db := db.GetDB()
	var co Coin

	if err := db.Where("name = ?", username).Delete(&co).Error; err != nil {
		return err
	}

	return nil
}
