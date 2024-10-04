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

type TransactionHandler struct {
	Service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		Service: service,
	}
}

func (h *TransactionHandler) BuyProduct(c *gin.Context) {
	var body request.BuyProduct
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	userUuid := c.GetString("uuid")
	if err := h.Service.BuyProduct(userUuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Produk berhasil diajukan pembelian"))
}

func (h *TransactionHandler) DeleteProductTransaction(c *gin.Context) {

	userUuid := c.GetString("uuid")
	uuid := c.Param("uuid")
	if err := h.Service.DeleteProductTransaction(userUuid, uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Transaksi produk berhasil dihapus"))
}

func (h *TransactionHandler) AcceptProductTransaction(c *gin.Context) {

	userUuid := c.GetString("uuid")
	transactionUuid := c.Param("transactionUuid")
	if err := h.Service.AcceptProductTransaction(userUuid, transactionUuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Transaksi produk berhasil diselesaikan"))
}

func (h *TransactionHandler) UnacceptProductTransaction(c *gin.Context) {

	userUuid := c.GetString("uuid")
	transactionUuid := c.Param("transactionUuid")
	if err := h.Service.UnacceptProductTransaction(userUuid, transactionUuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Transaksi produk berhasil ditolak"))
}

func (h *TransactionHandler) SetTransactionProof(c *gin.Context) {
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

	proofFolder := "internal/files/transaction_proofs"
	filename := utils.RandomFileName(file)
	pathFile := filepath.Join(proofFolder, filename)

	if err := utils.SaveUploadedFile(file, pathFile); err != nil {
		if err := os.Remove(pathFile); err != nil {
			log.Println(err.Error())
		}
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	userUuid := c.GetString("uuid")
	transactionUuid := c.Param("transactionUuid")
	if err := h.Service.SetTransactionProof(userUuid, transactionUuid, filename); err != nil {
		if err := os.Remove(pathFile); err != nil {
			log.Println(err.Error())
		}
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Bukti Transaksi produk berhasil diupload"))
}
