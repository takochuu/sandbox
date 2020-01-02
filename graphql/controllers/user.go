package controllers

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) GetUser() {
}