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
	RoleRepository repository.RoleRepository
	PasswordUtils  helper.PasswordUtils
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(
	roleRepository repository.RoleRepository,
	userRepository repository.UserRepository,
	passwordUtils helper.PasswordUtils,
	validate *validator.Validate) *UserServiceImpl {
	userService := &UserServiceImpl{
		RoleRepository: roleRepository,
		UserRepository: userRepository,
		Validate:       validate,
		PasswordUtils:  passwordUtils,
	}
	return userService
}

func (service *UserServiceImpl) Save(ctx *gin.Context) {

	var request web.UserCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(errors.NewLudesError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := service.Validate.Struct(request); err != nil {
		ctx.Error(errors.NewLudesError(http.StatusBadRequest, err.Error()))
		return
	}
	hashedPassword, err := service.PasswordUtils.HashPassword(request.Password)
	if err != nil {
		ctx.Error(errors.NewLudesError(http.StatusInternalServerError, err.Error()))
		return
	}
	var roles []domain.Role
	for _, roleRequest := range request.Roles {
		role, errRole := service.RoleRepository.FindByID(roleRequest.RoleID)
		if errRole != nil {
			ctx.Error(errors.NewLudesError(http.StatusBadRequest, "Role not found"))
			return
		}
		roles = append(roles, *role)
	}
	user := &domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
		Username: request.Username,
		Roles:    roles,
	}
	user, err = service.UserRepository.Save(user)
	if err != nil {
		ctx.Error(errors.NewLudesError(http.StatusInternalServerError, err.Error()))
		return
	}

	response := web.Response(http.StatusCreated, "Success", user)
	ctx.JSON(http.StatusCreated, response)

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

func (service *UserServiceImpl) Update(ctx *gin.Context) {

	userId := helper.ConvertStringToInt(ctx.Param("userId"))
	request := web.UserUpdateRequest{
		ID: uint(userId),
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(errors.NewLudesError(http.StatusBadRequest, err.Error()))
		return
	}
	if err := service.Validate.Struct(request); err != nil {
		ctx.Error(errors.NewLudesError(http.StatusBadRequest, err.Error()))
		return
	}

	user, err := service.UserRepository.FindByID(uint(userId))
	if err != nil {
		ctx.Error(errors.NewLudesError(http.StatusNotFound, "Not found"))
		return
	}
	service.UpdateUserFields(user, request)
	user, err = service.UserRepository.Update(user)
	if err != nil {
		ctx.Error(errors.NewLudesError(http.StatusInternalServerError, err.Error()))
		return
	}

	response := web.Response(http.StatusOK, "Success", user)
	ctx.JSON(http.StatusOK, response)
}

func (service *UserServiceImpl) Delete(ctx *gin.Context) {
	userId := helper.ConvertStringToInt(ctx.Param("userId"))
	user, err := service.UserRepository.FindByID(uint(userId))
	if err != nil {
		ctx.Error(errors.NewLudesError(http.StatusNotFound, "Not found"))
		return
	}
	service.UserRepository.Delete(user)

	response := web.Response(http.StatusOK, "Success", web.Null())
	ctx.JSON(http.StatusOK, response)
}

func (service *UserServiceImpl) FindByID(ctx *gin.Context) {
	userId := helper.ConvertStringToInt(ctx.Param("userId"))
	user, err := service.UserRepository.FindByID(uint(userId))

	if err != nil {
		ctx.Error(errors.NewLudesError(http.StatusNotFound, "Not found"))
		return
	}

	response := web.Response(http.StatusOK, "Success", web.NewUserResponse(*user))
	ctx.JSON(http.StatusOK, response)
}

func (service *UserServiceImpl) FindAll(ctx *gin.Context) {
	users, _ := service.UserRepository.FindAll()
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, web.NewUserResponse(user))
	}
	response := web.Response(http.StatusOK, "Success", userResponses)

	ctx.JSON(http.StatusOK, response)
}

func (service *UserServiceImpl) FindAllUserRoles(ctx *gin.Context) {
	roles, _ := service.RoleRepository.FindAll()
	var roleResponse []web.UserRoleResponse
	for _, role := range roles {
		roleResponse = append(roleResponse, web.UserRoleResponse{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
		})
	}
	response := web.Response(http.StatusOK, "Success", roleResponse)
	ctx.JSON(http.StatusOK, response)
}
