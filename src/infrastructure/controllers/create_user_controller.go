package controllers

import (
	"Recu_ArqSoftware/src/application/usecases"
	_ "Recu_ArqSoftware/src/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserController struct {
	createUserUseCase *usecases.CreateUserUseCase
}

func NewCreateUserController(createUserUseCase *usecases.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{createUserUseCase: createUserUseCase}
}

func (c *CreateUserController) Handle(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.createUserUseCase.Execute(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
