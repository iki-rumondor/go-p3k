package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-p3k/internal/app/layers/services"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/request"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/response"
	"github.com/iki-rumondor/go-p3k/internal/utils"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

func (h *AuthHandler) RegisterGuest(c *gin.Context) {
	var body request.RegisterGuest
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.RegisterGuest(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Akun berhasil didaftarkan"))
}

func (h *AuthHandler) RegisterShop(c *gin.Context) {
	var body request.RegisterShop
	if err := c.Bind(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	shopImage, err := c.FormFile("shop_image")
	if err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if status := utils.CheckTypeFile(shopImage, []string{"jpg", "png", "jpeg"}); !status {
		utils.HandleError(c, response.BADREQ_ERR("Tipe file pada gambar toko tidak valid, gunakan tipe jpg, png, atau jpeg"))
		return
	}

	if moreThan := utils.CheckFileSize(shopImage, 1); moreThan {
		utils.HandleError(c, response.BADREQ_ERR("Gambar toko yang diupload lebih dari 1MB"))
		return
	}

	shopsFolder := "internal/files/shops"
	shopFile := utils.RandomFileName(shopImage)
	shopPath := filepath.Join(shopsFolder, shopFile)

	if err := utils.SaveUploadedFile(shopImage, shopPath); err != nil {
		if err := os.Remove(shopPath); err != nil {
			log.Println(err.Error())
		}
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	identityImage, err := c.FormFile("identity_image")
	if err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if status := utils.CheckTypeFile(identityImage, []string{"jpg", "png", "jpeg"}); !status {
		utils.HandleError(c, response.BADREQ_ERR("Tipe file pada gambar ktp tidak valid, gunakan tipe jpg, png, atau jpeg"))
		return
	}

	if moreThan := utils.CheckFileSize(identityImage, 1); moreThan {
		utils.HandleError(c, response.BADREQ_ERR("Gambar ktp yang diupload lebih dari 1MB"))
		return
	}

	identitiesFolder := "internal/files/identities"
	identityFile := utils.RandomFileName(identityImage)
	identityPath := filepath.Join(identitiesFolder, identityFile)

	if err := utils.SaveUploadedFile(identityImage, identityPath); err != nil {
		if err := os.Remove(identityPath); err != nil {
			log.Println(err.Error())
		}
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	if err := h.Service.RegisterShop(shopFile, identityFile, &body); err != nil {
		if err := os.Remove(identityPath); err != nil {
			log.Println(err.Error())
		}
		if err := os.Remove(shopPath); err != nil {
			log.Println(err.Error())
		}
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Akun berhasil didaftarkan"))
}

func (h *AuthHandler) VerifyUser(c *gin.Context) {
	var body request.SignIn
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	resp, err := h.Service.VerifyUser(&body)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AuthHandler) GetUserByUuid(c *gin.Context) {

	uuid := c.GetString("uuid")
	if uuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	resp, err := h.Service.GetUserByUuid(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AuthHandler) ActivationUser(c *gin.Context) {
	var body request.Activation
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.ActivationUser(uuid, body.Status); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Status aktivasi berhasil diperbarui"))
}
