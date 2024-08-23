package handler

import (
	"fmt"
	model "shopifun/Model"
	"shopifun/helper"
	"shopifun/request"
	"shopifun/service"
	"shopifun/utils"

	"github.com/gofiber/fiber/v3"
)

type ResponseMessage struct {
	Status  int
	Message string
	Data    interface{}
}

type registerResponse struct {
	ID       uint
	Username string
	Email    string
}

func Register(c fiber.Ctx) error {
	registerUser := new(model.User)

	if err := c.Bind().JSON(registerUser); err != nil {
		return err
	}
	// validate email
	validateEmail := utils.IsEmail(registerUser.Email)

	if !validateEmail {
		sendData := Response{
			Status:  400,
			Message: "Invalid Email",
		}

		return c.Status(fiber.StatusBadRequest).JSON(sendData)
	}

	validatePassword := utils.CheckLenPassword(registerUser.Password)

	if !validatePassword {
		sendData := Response{
			Status:  400,
			Message: "Password must be 6 character or more",
		}

		return c.Status(fiber.StatusBadRequest).JSON(sendData)
	}

	hashedpassword, err := utils.HashPassword(registerUser.Password)

	if err != nil {
		fmt.Println("mantap")
	}

	registerUser.Password = hashedpassword

	resultInsertChan := make(chan *model.User)
	dbErrChan := make(chan error)

	go func() {
		resultInsert, err := service.InsertUser(registerUser)

		if err != nil {
			dbErrChan <- err
			return
		}
		resultInsertChan <- resultInsert
	}()

	select {
	case resultInsert := <-resultInsertChan:
		userResponse := registerResponse{
			ID:       resultInsert.ID,
			Username: resultInsert.Username,
			Email:    resultInsert.Email,
		}
		returnResponse := helper.ApiResponse("Success Create New User", fiber.StatusCreated, userResponse)

		return c.Status(fiber.StatusOK).JSON(returnResponse)

	case err := <-dbErrChan:
		returnResponse := helper.ApiResponse(err.Error(), fiber.StatusBadRequest, nil)
		return c.Status(fiber.StatusBadRequest).JSON(returnResponse)

	}
}

func Login(c fiber.Ctx) error {
	var loginUserRequest request.LoginUserRequest

	if err := c.Bind().JSON(&loginUserRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// check email exist or not
	checkUser, _ := service.GetUsername(loginUserRequest.Username)

	if checkUser == nil {
		apiReturn := helper.ApiResponse("User Not Found", 404, nil)

		return c.Status(fiber.StatusNotFound).JSON(apiReturn)
	}
	passwordDB := checkUser.Password

	comparePasswordDB := utils.ComparePassword(passwordDB, loginUserRequest.Password)

	if !comparePasswordDB {
		apiReturn := helper.ApiResponse("Email / Password Invalid", 403, nil)
		return c.Status(fiber.StatusForbidden).JSON(apiReturn)
	}

	tokenChan := make(chan string)
	tokenErrChan := make(chan error)

	// get jwt token

	go func() {
		t, err := utils.GenerateJWTToken(loginUserRequest.Username, checkUser.ID)

		if err != nil {
			tokenErrChan <- err
			return
		}

		tokenChan <- t
	}()

	select {
	case t := <-tokenChan:
		returnData := helper.ApiResponse("Success Login", 200, t)
		return c.Status(fiber.StatusOK).JSON(returnData)

	case err := <-tokenErrChan:
		fmt.Println("error", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

}

func DetailProfile(c fiber.Ctx) error {
	fmt.Println("masuk sini")
	return nil
}
