package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"login-api/service"
)

type Controller struct{}

// Index action: GET
func (pc Controller) Index(c *gin.Context) {
	var s service.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST
func (pc Controller) Create(c *gin.Context) {
	var s service.Service
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Showid action: GET
func (pc Controller) Showid(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Showusername action: GET
func (pc Controller) Showname(c *gin.Context) {
	username := c.Params.ByName("username")
	var s service.Service
	p, err := s.GetByName(username)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		c.JSON(200, err)
	} else {
		fmt.Println("ss")
		c.JSON(200, p)
	}
}

// Update action: PUT
func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE
func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}

func (pc Controller) MemoIndex(c *gin.Context) {
	var s service.Service
	p, err := s.MemoGetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) MemoCreate(c *gin.Context) {
	var s service.Service
	p, err := s.MemoCreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

func (pc Controller) MemoDelete(c *gin.Context) {
	username := c.Params.ByName("username")
	var s service.Service

	if err := s.MemoDeleteByID(username); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + username: "deleted"})
	}
}

func (pc Controller) MemoShowname(c *gin.Context) {
	username := c.Params.ByName("username")
	var s service.Service
	p, err := s.MemoGetByName(username)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		c.JSON(200, err)
	} else {
		fmt.Println(p)
		c.JSON(200, p)
	}
}

func (pc Controller) TicketAdd(c *gin.Context) {
	username := c.Params.ByName("username")
	var s service.Service
	s.TicketAdd(username)
}

// func (pc Controller) CoinAdd(c *gin.Context) {
// 	username := c.Params.ByName("username")
// 	var s service.Service
// 	s.CoinAdd(username)
// }

func (pc Controller) TicketSub(c *gin.Context) {
	username := c.Params.ByName("username")
	var s service.Service
	s.TicketSub(username)
}

// func (pc Controller) CoinSub(c *gin.Context) {
// 	username := c.Params.ByName("username")
// 	var s service.Service
// 	s.CoinSub(username)
// }

func (pc Controller) GameCoinAll(c *gin.Context) {
	var s service.Service
	p, err := s.GameCoinGetAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) GameCoinCreate(c *gin.Context) {
	var s service.Service
	p, err := s.GameCoinCreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

func (pc Controller) GameCoinShowname(c *gin.Context) {
	username := c.Params.ByName("username")
	var s service.Service
	p, err := s.GameCoinGetByName(username)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		c.JSON(200, err)
	} else {
		fmt.Println(p)
		c.JSON(200, p)
	}
}

func (pc Controller) GameCoinUpdate(c *gin.Context) {
	username := c.Params.ByName("username")
	var s service.Service
	p, err := s.GameCoinUpdateByName(username, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) GameCoinDelete(c *gin.Context) {
	username := c.Params.ByName("username")
	var s service.Service

	if err := s.GameCoinDeleteByName(username); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + username: "deleted"})
	}
}
