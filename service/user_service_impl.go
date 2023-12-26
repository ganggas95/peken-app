package service

import (
	"net/http"
	"peken-be/helper"
	"peken-be/models/domain"
	"peken-be/models/errors"
	"peken-be/models/web"
	"peken-be/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserServiceImpl struct {
	PasswordUtils  helper.PasswordUtils
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(
	userRepository repository.UserRepository,
	passwordUtils helper.PasswordUtils,
	validate *validator.Validate) *UserServiceImpl {
	userService := &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
		PasswordUtils:  passwordUtils,
	}
	return userService
}

func (service *UserServiceImpl) Save(ctx *gin.Context) *domain.User {

	var request web.UserCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return service.HandleError(ctx, http.StatusBadRequest, err)
	}
	if err := service.Validate.Struct(request); err != nil {
		return service.HandleError(ctx, http.StatusBadRequest, err)
	}
	hashedPassword, err := service.PasswordUtils.HashPassword(request.Password)
	if err != nil {
		return service.HandleError(ctx, http.StatusInternalServerError, err)
	}
	user := &domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
		Username: request.Username,
	}
	user, err = service.UserRepository.Save(user)
	if err != nil {
		return service.HandleError(ctx, http.StatusInternalServerError, err)
	}
	return user
}

func (service *UserServiceImpl) HandleError(ctx *gin.Context, status int, err error) *domain.User {
	cError := errors.NewLudesError(http.StatusBadRequest, err.Error())
	ctx.Error(cError)
	return nil
}

func (service *UserServiceImpl) UpdateUserFields(user *domain.User, request web.UserUpdateRequest) {
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Username != "" {
		user.Username = request.Username
	}
}

func (service *UserServiceImpl) Update(ctx *gin.Context) *domain.User {

	userId := helper.ConvertStringToInt(ctx.Param("userId"))
	request := web.UserUpdateRequest{
		ID: uint(userId),
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return service.HandleError(ctx, http.StatusBadRequest, err)
	}
	if err := service.Validate.Struct(request); err != nil {
		return service.HandleError(ctx, http.StatusBadRequest, err)
	}

	user, err := service.UserRepository.FindByID(uint(userId))
	if err != nil {
		return service.HandleError(ctx, http.StatusNotFound, err)
	}
	service.UpdateUserFields(user, request)
	user, err = service.UserRepository.Update(user)
	if err != nil {
		return service.HandleError(ctx, http.StatusInternalServerError, err)
	}
	return user
}

func (service *UserServiceImpl) Delete(ctx *gin.Context) *int {
	userId := helper.ConvertStringToInt(ctx.Param("userId"))
	user, err := service.UserRepository.FindByID(uint(userId))
	if err != nil {
		service.HandleError(ctx, http.StatusNotFound, err)
		return nil
	}
	service.UserRepository.Delete(user)
	return &userId
}

func (service *UserServiceImpl) FindByID(ctx *gin.Context) *domain.User {
	userId := helper.ConvertStringToInt(ctx.Param("userId"))
	user, err := service.UserRepository.FindByID(uint(userId))

	if err != nil {
		return service.HandleError(ctx, http.StatusNotFound, err)
	}

	return user
}

func (service *UserServiceImpl) FindAll(ctx *gin.Context) []domain.User {
	users, _ := service.UserRepository.FindAll()
	return users
}
