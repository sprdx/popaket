package users

import (
	"net/http"
	"popaket/businesses/users"
	requests "popaket/controllers/users/request"
	responses "popaket/controllers/users/response"
	"popaket/helpers"
	res "popaket/responses"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Usecase users.Usecase
}

func NewUserController(u users.Usecase) *UserController {
	return &UserController{
		Usecase: u,
	}
}

func (u *UserController) Register(c echo.Context) error {
	userRegister := requests.RegisterRequest{}
	c.Bind(&userRegister)
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	ctx := c.Request().Context()

	userDomain := users.Domain{
		Id:       uuid,
		Name:     userRegister.Name,
		Username: userRegister.Username,
		Password: userRegister.Password,
		MSISDN:   userRegister.MSISDN,
	}

	user, err := u.Usecase.Register(ctx, userDomain)
	if err != nil {
		return c.JSON(http.StatusBadRequest, res.BadRequestResponse("Bad request"))
	}

	registerResponse := responses.UserResponse{
		ID:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		MSISDN:   user.MSISDN,
	}

	return c.JSON(http.StatusOK, res.SuccessResponseData("Congratulation! User created successfully", registerResponse))
}

func (u *UserController) Login(c echo.Context) error {
	userLogin := requests.LoginRequest{}
	c.Bind(&userLogin)

	ctx := c.Request().Context()

	userDomain := users.Domain{
		MSISDN:   userLogin.MSISDN,
		Password: userLogin.Password,
	}

	user, err := u.Usecase.Login(ctx, userDomain.MSISDN, userDomain.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, res.BadRequestResponse("MSISDN or password is incorrect"))
	}

	loginResponse := responses.LoginResponse{
		Token: user.Token,
	}

	return c.JSON(http.StatusOK, res.SuccessResponseData("Login success", loginResponse))
}

func (u *UserController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	userId := helpers.ExtractTokenUserId(c)

	user, err := u.Usecase.GetById(ctx, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, res.BadRequestResponse("Operation Failed"))
	}

	response := responses.UserResponse{
		ID:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		MSISDN:   user.MSISDN,
	}

	return c.JSON(http.StatusOK, res.SuccessResponseData("Operation success", response))
}
