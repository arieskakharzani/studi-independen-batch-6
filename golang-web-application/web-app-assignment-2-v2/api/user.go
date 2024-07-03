package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	// TODO: answer here
	var user model.User

	// Binding data dari body request ke struct User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid decode json"})
		return
	}

	// Memeriksa apakah email atau password kosong
	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is empty"})
		return
	}

	// Melakukan login dengan memanggil service
	token, err := u.userService.Login(&user)
	if err != nil {
		// Jika terjadi error saat login, kembalikan response HTTP 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error internal server"})
		return
	}

	// Jika login berhasil, membuat cookie session_token
	c.SetCookie("session_token", *token, 3600, "/", "localhost", false, true)

	// Mengirim response dengan status code 200 dan data user yang sudah login
	c.JSON(http.StatusOK, gin.H{"user_id": user.ID, "message": "login success"})
}

func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	// TODO: answer here
	// Mengambil daftar kategori tugas pengguna dari service
	userTaskCategories, err := u.userService.GetUserTaskCategory()
	if err != nil {
		// Jika terjadi error saat mengambil data, kembalikan response HTTP 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error internal server"})
		return
	}

	// Mengirim response dengan status code 200 dan daftar kategori tugas pengguna
	c.JSON(http.StatusOK, userTaskCategories)
}
