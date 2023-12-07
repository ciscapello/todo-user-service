package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"todo-user-service/internal/app/utils"
)

type registrationRequestBody struct {
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (h AuthHandler) Registration(writer http.ResponseWriter, req *http.Request) {
	var body registrationRequestBody
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		message := "Bad request"
		utils.ReturnBadRequestError(writer, http.StatusBadRequest, "message", message)
		return
	}
	err = validateRegistrationBody(writer, body)
	if err != nil {
		return
	}

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		message := "Cannot hash password"
		utils.ReturnBadRequestError(writer, http.StatusBadRequest, "message", message)
		return
	}

	id, err := h.userService.CreateUser(body.Phone, hashedPassword)
	if err != nil {
		utils.ReturnBadRequestError(writer, http.StatusBadRequest, "message", err.Error())
	}

	writer.WriteHeader(http.StatusOK)
	resp := make(map[string]interface{})
	resp["message"] = "ok"
	resp["id"] = id
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	_, _ = writer.Write(jsonResp)
}

func validateRegistrationBody(writer http.ResponseWriter, body registrationRequestBody) error {
	if len(body.Phone) < 11 {
		message := "phone number must be at least 11 symbols"
		utils.ReturnBadRequestError(writer, http.StatusBadRequest, "message", message)
		return errors.New(message)
	}

	if len(body.Password) < 8 {
		message := "password number must be at least 8 symbols"
		utils.ReturnBadRequestError(writer, http.StatusBadRequest, "message", message)
		return errors.New(message)
	}

	if body.Password != body.ConfirmPassword {
		message := "password should be equal to confirm password"
		utils.ReturnBadRequestError(writer, http.StatusBadRequest, "message", message)
		return errors.New(message)
	}
	return nil
}
