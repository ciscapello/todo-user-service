package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"todo-user-service/internal/app/utils"

	"github.com/gin-gonic/gin"
)

type loginRequestBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var body loginRequestBody
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		return
	}
	err = validateBody(c.Writer, body)
	if err != nil {
		return
	}

	accessToken, err := h.tokenManager.GenerateAccessToken(body.Username)
	if err != nil {
		fmt.Println(err)
	}
	refreshToken, err := h.tokenManager.GenerateRefreshToken(body.Username)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "ok",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})

}

func validateBody(writer http.ResponseWriter, body loginRequestBody) error {
	if len(body.Username) < 5 {
		message := "username must be at least 11 symbols"
		utils.ReturnBadRequestError(writer, http.StatusBadRequest, "message", message)
		return errors.New(message)
	}

	if len(body.Password) < 8 {
		message := "password number must be at least 8 symbols"
		utils.ReturnBadRequestError(writer, http.StatusBadRequest, "message", message)
		return errors.New(message)
	}

	return nil
}
