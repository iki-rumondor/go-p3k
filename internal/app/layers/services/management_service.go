package services

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/iki-rumondor/go-p3k/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/request"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/response"
	"github.com/iki-rumondor/go-p3k/internal/utils"
	"gorm.io/gorm"
)

type ManagementService struct {
	Repo interfaces.ManagementInterface
}

func NewManagementService(repo interfaces.ManagementInterface) *ManagementService {
	return &ManagementService{
		Repo: repo,
	}
}

func (s *ManagementService) CreateCategory(req *request.Category) error {
	model := models.Category{
		Name: req.Name,
	}

	if err := s.Repo.CreateModel(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) UpdateCategory(uuid string, req *request.Category) error {

	model := models.Category{
		Name: req.Name,
	}

	if err := s.Repo.UpdateCategory(uuid, &model); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Kategori tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) CreateShop(req *request.Shop) error {
	var username = utils.GenerateRandomString(8)

	model := models.Shop{
		Name:        req.Name,
		Owner:       req.Owner,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		User: &models.User{
			Name:     req.Owner,
			Username: username,
			Password: username,
			RoleID:   3,
			Active:   true,
		},
	}

	if err := s.Repo.CreateShop(req.CategoryUuid, &model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) UpdateShop(uuid string, req *request.Shop) error {

	model := models.Shop{
		Name:        req.Name,
		Owner:       req.Owner,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
	}

	if err := s.Repo.UpdateShop(uuid, req.CategoryUuid, &model); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Umkm tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) CreateProduct(userUuid, imageName string, req *request.Product) error {
	price, err := strconv.Atoi(req.Price)
	if err != nil {
		return response.BADREQ_ERR("Harga yang dimasukkan tidak valid")
	}

	stock, err := strconv.Atoi(req.Stock)
	if err != nil {
		return response.BADREQ_ERR("Stok yang dimasukkan tidak valid")
	}

	model := models.Product{
		Name:  req.Name,
		Price: int64(price),
		Stock: int64(stock),
		Image: imageName,
	}

	if err := s.Repo.CreateProduct(userUuid, &model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) UpdateProduct(userUuid, uuid, imageName string, req *request.Product) error {
	price, err := strconv.Atoi(req.Price)
	if err != nil {
		return response.BADREQ_ERR("Harga yang dimasukkan tidak valid")
	}

	stock, err := strconv.Atoi(req.Stock)
	if err != nil {
		return response.BADREQ_ERR("Stok yang dimasukkan tidak valid")
	}

	model := models.Product{
		Name:  req.Name,
		Price: int64(price),
		Stock: int64(stock),
		Image: imageName,
	}

	oldImage, err := s.Repo.UpdateProduct(userUuid, uuid, &model)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Produk tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	productsFolder := "internal/files/products"
	pathFile := filepath.Join(productsFolder, oldImage)

	if imageName != "" {
		if err := os.Remove(pathFile); err != nil {
			log.Println(err.Error())
		}
	}

	return nil
}

// func (s *ManagementService) DeleteMajor(uuid string) error {

// 	if err := s.Repo.DeleteMajor(uuid); err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return response.NOTFOUND_ERR("Jurusan tidak ditemukan")
// 		}
// 		return response.SERVICE_INTERR
// 	}

// 	return nil
// }
