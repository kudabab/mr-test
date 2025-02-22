package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kudabab/mr-test/entity"
	"github.com/kudabab/mr-test/service"
)

type UserController interface {
	Save(ctx *gin.Context) entity.User
	FindAll() []entity.User
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}
func (c *controller) FindAll() []entity.User {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entity.User {

	var user entity.User
	ctx.BindJSON(&user)
	c.service.Save(user)
	return user
}
