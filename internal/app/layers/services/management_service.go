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

// func (s *ManagementService) CreateShop(req *request.Shop) error {
// 	var username = utils.GenerateRandomString(8)

// 	model := models.Shop{
// 		Name:        req.Name,
// 		Owner:       req.Owner,
// 		Address:     req.Address,
// 		PhoneNumber: req.PhoneNumber,
// 		User: &models.User{
// 			Name:     req.Owner,
// 			Username: username,
// 			Password: username,
// 			RoleID:   3,
// 			Active:   true,
// 		},
// 	}

// 	if err := s.Repo.CreateShop(req.CategoryUuid, &model); err != nil {
// 		return response.SERVICE_INTERR
// 	}

// 	return nil
// }

// func (s *ManagementService) UpdateShop(uuid string, req *request.Shop) error {

// 	model := models.Shop{
// 		Name:        req.Name,
// 		Owner:       req.Owner,
// 		Address:     req.Address,
// 		PhoneNumber: req.PhoneNumber,
// 	}

// 	if err := s.Repo.UpdateShop(uuid, req.CategoryUuid, &model); err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return response.NOTFOUND_ERR("Umkm tidak ditemukan")
// 		}
// 		return response.SERVICE_INTERR
// 	}

// 	return nil
// }

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
		Unit:  req.Unit,
		Price: int64(price),
		Stock: int64(stock),
		Image: imageName,
	}

	if err := s.Repo.CreateProduct(userUuid, req.CategoryUuid, &model); err != nil {
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
		Unit:  req.Unit,
		Price: int64(price),
		Stock: int64(stock),
		Image: imageName,
	}

	oldImage, err := s.Repo.UpdateProduct(userUuid, req.CategoryUuid, uuid, &model)
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

func (s *ManagementService) CreateCitizen(req *request.Citizen) error {
	if unique := s.Repo.CheckUniqueNik(req.Nik); !unique {
		return response.BADREQ_ERR("Nik yang digunakan sudah terdafatar")
	}
	model := models.Citizen{
		Name:        req.Name,
		Nik:         req.Nik,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		User: &models.User{
			Name:     req.Name,
			Username: req.Nik,
			Password: req.Nik,
			RoleID:   5,
			Active:   true,
		},
	}

	if err := s.Repo.CreateModel(&model); err != nil {
		log.Println(err)
		if utils.IsErrorType(err) {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) UpdateCitizen(uuid string, req *request.Citizen) error {

	model := models.Citizen{
		Name:        req.Name,
		Nik:         req.Nik,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
	}

	if err := s.Repo.UpdateCitizen(uuid, &model); err != nil {
		log.Println(err)
		if utils.IsErrorType(err) {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) CreateMember(req *request.Member) error {
	var username = utils.GenerateRandomString(8)

	model := models.Member{
		Name:        req.Name,
		IsImportant: req.IsImportant,
		Position:    req.Position,
		User: &models.User{
			Name:     req.Name,
			Username: username,
			Password: username,
			RoleID:   2,
			Active:   true,
		},
	}

	if err := s.Repo.CreateModel(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) UpdateMember(uuid string, req *request.Member) error {

	model := models.Member{
		Name:        req.Name,
		IsImportant: req.IsImportant,
		Position:    req.Position,
	}

	if err := s.Repo.UpdateMember(uuid, &model); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Anggota tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) CreateActivity(userUuid, imageName string, req *request.Activity) error {
	user, err := s.Repo.GetUserByUuid(userUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("User tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}
	if user.RoleID != 1 && user.RoleID != 2 {
		return response.UNAUTH_ERR("Akses dibatasi")
	}

	model := models.Activity{
		Title:       req.Title,
		Description: req.Description,
		Group:       req.Group,
		ImageName:   imageName,
	}

	if err := s.Repo.CreateActivity(userUuid, &model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) UpdateActivity(userUuid, uuid, imageName string, req *request.Activity) error {
	user, err := s.Repo.GetUserByUuid(userUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("User tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}
	if user.RoleID != 1 && user.RoleID != 2 {
		return response.UNAUTH_ERR("Akses dibatasi")
	}

	model := models.Activity{
		Title:       req.Title,
		Description: req.Description,
		Group:       req.Group,
		ImageName:   imageName,
	}

	oldImage, err := s.Repo.UpdateActivity(userUuid, uuid, &model)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Kegiatan tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	activitiesFolder := "internal/files/activities"
	pathFile := filepath.Join(activitiesFolder, oldImage)

	if imageName != "" {
		if err := os.Remove(pathFile); err != nil {
			log.Println(err.Error())
		}
	}

	return nil
}

func (s *ManagementService) DeleteActivity(userUuid, uuid string) error {
	user, err := s.Repo.GetUserByUuid(userUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("User tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}
	if user.RoleID != 1 && user.RoleID != 2 {
		return response.UNAUTH_ERR("Akses dibatasi")
	}

	imageName, err := s.Repo.DeleteActivity(uuid)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Kegiatan tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	activitiesFolder := "internal/files/activities"
	pathFile := filepath.Join(activitiesFolder, imageName)
	if err := os.Remove(pathFile); err != nil {
		log.Println(err.Error())
	}

	return nil
}

func (s *ManagementService) CreateMemberActivity(userUuid string, req *request.MemberActivity) error {
	user, err := s.Repo.GetUserByUuid(userUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("User tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}
	if user.RoleID != 1 && user.RoleID != 2 {
		return response.UNAUTH_ERR("Akses dibatasi")
	}

	exist, err := s.Repo.CheckExistMemberActivity(req.MemberUuid, req.ActivityUuid)
	if err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	if exist {
		return response.BADREQ_ERR("Anggota telah ditambahkan")
	}

	if err := s.Repo.CreateMemberActivity(user.ID, req.MemberUuid, req.ActivityUuid); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *ManagementService) DeleteMemberActivity(userUuid, memberUuid, activityUuid string) error {
	user, err := s.Repo.GetUserByUuid(userUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("User tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}
	if user.RoleID != 1 && user.RoleID != 2 {
		return response.UNAUTH_ERR("Akses dibatasi")
	}

	if err := s.Repo.DeleteMemberActivity(memberUuid, activityUuid); err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Anggota kegiatan tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	return nil
}
