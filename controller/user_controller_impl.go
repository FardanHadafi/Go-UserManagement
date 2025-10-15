package controller

import (
	"Go-UserManagement/helper"
	"Go-UserManagement/model/web"
	"Go-UserManagement/service"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewCategoryController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (Controller *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var request web.UserRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	userResponse, err := Controller.UserService.Register(r.Context(), request)
	if err != nil {
		helper.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.WriteSuccessResponse(w, http.StatusCreated, userResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)

	userResponse, err := controller.UserService.Login(request.Context(), userLoginRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}