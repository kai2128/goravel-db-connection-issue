package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Show(ctx http.Context) http.Response {
	var user1s []models.User
	var user2s []models.User2
	facades.Orm().Query().Get(&user1s)
	facades.Orm().Query().Get(&user2s)
	return ctx.Response().Success().Json(http.Json{
		"user1s": user1s,
		"user2s": user2s,
	})
}
