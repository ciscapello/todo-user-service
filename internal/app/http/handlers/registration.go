package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"todo-user-service/internal/app/utils"
)

type requestBody struct {
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func validateRegistrationBody(writer http.ResponseWriter, body requestBody) error {
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

func (h Handler) Registration(writer http.ResponseWriter, req *http.Request) {
	var body requestBody
	err := json.NewDecoder(req.Body).Decode(&body)
	err = validateRegistrationBody(writer, body)
	if err != nil {
		return
	}
	writer.WriteHeader(http.StatusOK)
	resp := make(map[string]string)
	resp["message"] = "ok"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	_, _ = writer.Write(jsonResp)

	return
}
