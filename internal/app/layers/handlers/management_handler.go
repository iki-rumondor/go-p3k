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

type ManagementHandler struct {
	Service *services.ManagementService
}

func NewManagementHandler(service *services.ManagementService) *ManagementHandler {
	return &ManagementHandler{
		Service: service,
	}
}

func (h *ManagementHandler) CreateCategory(c *gin.Context) {
	var body request.Category
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreateCategory(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Kategori Berhasil Ditambahkan"))
}

func (h *ManagementHandler) UpdateCategory(c *gin.Context) {
	var body request.Category
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateCategory(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Kategori Berhasil Diperbarui"))
}

func (h *ManagementHandler) CreateShop(c *gin.Context) {
	var body request.Shop
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreateShop(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Umkm Berhasil Ditambahkan"))
}

func (h *ManagementHandler) UpdateShop(c *gin.Context) {
	var body request.Shop
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateShop(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Umkm Berhasil Diperbarui"))
}

func (h *ManagementHandler) CreateProduct(c *gin.Context) {
	var body request.Product
	if err := c.Bind(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if status := utils.CheckTypeFile(file, []string{"jpg", "png", "jpeg"}); !status {
		utils.HandleError(c, response.BADREQ_ERR("Tipe file tidak valid, gunakan tipe jpg, png, atau jpeg"))
		return
	}

	if moreThan := utils.CheckFileSize(file, 1); moreThan {
		utils.HandleError(c, response.BADREQ_ERR("File yang diupload lebih dari 1MB"))
		return
	}

	productsFolder := "internal/files/products"
	filename := utils.RandomFileName(file)
	pathFile := filepath.Join(productsFolder, filename)

	if err := utils.SaveUploadedFile(file, pathFile); err != nil {
		if err := os.Remove(pathFile); err != nil {
			log.Println(err.Error())
		}
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	userUuid := c.GetString("uuid")
	if err := h.Service.CreateProduct(userUuid, filename, &body); err != nil {
		if err := os.Remove(pathFile); err != nil {
			log.Println(err.Error())
		}
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Produk Berhasil Ditambahkan"))
}

func (h *ManagementHandler) UpdateProduct(c *gin.Context) {
	var body request.Product
	if err := c.Bind(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}
	var filename string
	file, _ := c.FormFile("file")
	if file != nil {
		if status := utils.CheckTypeFile(file, []string{"jpg", "png", "jpeg"}); !status {
			utils.HandleError(c, response.BADREQ_ERR("Tipe file tidak valid, gunakan tipe jpg, png, atau jpeg"))
			return
		}

		if moreThan := utils.CheckFileSize(file, 1); moreThan {
			utils.HandleError(c, response.BADREQ_ERR("File yang diupload lebih dari 1MB"))
			return
		}

		productsFolder := "internal/files/products"
		filename = utils.RandomFileName(file)
		pathFile := filepath.Join(productsFolder, filename)

		if err := utils.SaveUploadedFile(file, pathFile); err != nil {
			if err := os.Remove(pathFile); err != nil {
				log.Println(err.Error())
			}
			utils.HandleError(c, response.HANDLER_INTERR)
			return
		}
	}

	userUuid := c.GetString("uuid")
	uuid := c.Param("uuid")
	if err := h.Service.UpdateProduct(userUuid, uuid, filename, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Produk Berhasil Diperbarui"))
}

func (h *ManagementHandler) CreateCitizen(c *gin.Context) {
	var body request.Citizen
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreateCitizen(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Data Masyarakat Berhasil Ditambahkan"))
}

func (h *ManagementHandler) UpdateCitizen(c *gin.Context) {
	var body request.Citizen
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateCitizen(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Data Masyarakat Berhasil Diperbarui"))
}

func (h *ManagementHandler) CreateMember(c *gin.Context) {
	var body request.Member
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreateMember(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Anggota Berhasil Ditambahkan"))
}

func (h *ManagementHandler) UpdateMember(c *gin.Context) {
	var body request.Member
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateMember(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Data Anggota Berhasil Diperbarui"))
}

func (h *ManagementHandler) CreateActivity(c *gin.Context) {
	var body request.Activity
	if err := c.Bind(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if status := utils.CheckTypeFile(file, []string{"jpg", "png", "jpeg"}); !status {
		utils.HandleError(c, response.BADREQ_ERR("Tipe file tidak valid, gunakan tipe jpg, png, atau jpeg"))
		return
	}

	if moreThan := utils.CheckFileSize(file, 1); moreThan {
		utils.HandleError(c, response.BADREQ_ERR("File yang diupload lebih dari 1MB"))
		return
	}

	activitiesFolder := "internal/files/activities"
	filename := utils.RandomFileName(file)
	pathFile := filepath.Join(activitiesFolder, filename)

	if err := utils.SaveUploadedFile(file, pathFile); err != nil {
		if err := os.Remove(pathFile); err != nil {
			log.Println(err.Error())
		}
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	userUuid := c.GetString("uuid")
	if err := h.Service.CreateActivity(userUuid, filename, &body); err != nil {
		if err := os.Remove(pathFile); err != nil {
			log.Println(err.Error())
		}
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Kegiatan Berhasil Ditambahkan"))
}

func (h *ManagementHandler) UpdateActivity(c *gin.Context) {
	var body request.Activity
	if err := c.Bind(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}
	var filename string
	file, _ := c.FormFile("file")
	if file != nil {
		if status := utils.CheckTypeFile(file, []string{"jpg", "png", "jpeg"}); !status {
			utils.HandleError(c, response.BADREQ_ERR("Tipe file tidak valid, gunakan tipe jpg, png, atau jpeg"))
			return
		}

		if moreThan := utils.CheckFileSize(file, 1); moreThan {
			utils.HandleError(c, response.BADREQ_ERR("File yang diupload lebih dari 1MB"))
			return
		}

		activitiesFolder := "internal/files/activities"
		filename = utils.RandomFileName(file)
		pathFile := filepath.Join(activitiesFolder, filename)

		if err := utils.SaveUploadedFile(file, pathFile); err != nil {
			if err := os.Remove(pathFile); err != nil {
				log.Println(err.Error())
			}
			utils.HandleError(c, response.HANDLER_INTERR)
			return
		}
	}

	userUuid := c.GetString("uuid")
	uuid := c.Param("uuid")
	if err := h.Service.UpdateActivity(userUuid, uuid, filename, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Kegiatan Berhasil Diperbarui"))
}

func (h *ManagementHandler) DeleteActivity(c *gin.Context) {
	userUuid := c.GetString("uuid")
	uuid := c.Param("uuid")
	if err := h.Service.DeleteActivity(userUuid, uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Kegiatan Berhasil Dihapus"))
}