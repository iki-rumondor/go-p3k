package services

import (
	"errors"
	"log"

	"github.com/iki-rumondor/go-p3k/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/request"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/response"
	"github.com/iki-rumondor/go-p3k/internal/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	Repo interfaces.AuthInterface
}

func NewAuthService(repo interfaces.AuthInterface) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}

func (s *AuthService) RegisterGuest(req *request.RegisterGuest) error {
	if req.Password != req.ConfirmPassword {
		return response.BADREQ_ERR("Konfirmasi password tidak sama dengan password")
	}

	if req.RoleID != 4 {
		return response.BADREQ_ERR("Role user tidak valid")
	}

	model := models.Guest{
		Name:        req.Fullname,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		User: &models.User{
			Name:     req.Fullname,
			Username: req.Username,
			Password: req.Password,
			Active:   false,
			RoleID:   req.RoleID,
		},
	}

	if err := s.Repo.CreateModel(&model); err != nil {
		log.Println(err.Error())
		if utils.IsErrorType(err) {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil

}

func (s *AuthService) RegisterShop(shopFile, identityFile string, req *request.RegisterShop) error {
	if req.Password != req.ConfirmPassword {
		return response.BADREQ_ERR("Konfirmasi password tidak sama dengan password")
	}

	if req.RoleID != 3 {
		return response.BADREQ_ERR("Role user tidak valid")
	}

	model := models.Shop{
		Name:          req.ShopName,
		Address:       req.Address,
		PhoneNumber:   req.PhoneNumber,
		Owner:         req.Owner,
		ShopImage:     shopFile,
		IdentityImage: identityFile,
		User: &models.User{
			Name:     req.Owner,
			Username: req.Username,
			Password: req.Password,
			Active:   false,
			RoleID:   req.RoleID,
		},
	}

	if err := s.Repo.CreateModel(&model); err != nil {
		log.Println(err.Error())
		if utils.IsErrorType(err) {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil

}

func (s *AuthService) VerifyUser(req *request.SignIn) (map[string]string, error) {
	user, err := s.Repo.FirstUserByUsername(req.Username)
	if err != nil {
		log.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NOTFOUND_ERR("Username atau Password Salah")
		}
		return nil, response.SERVICE_INTERR
	}

	if err := utils.ComparePassword(user.Password, req.Password); err != nil {
		return nil, response.NOTFOUND_ERR("Username atau Password Salah")
	}

	if !user.Active {
		return nil, response.NOTFOUND_ERR("Akun Anda Belum Diaktifkan")
	}

	jwt, err := utils.GenerateToken(user.Uuid, user.Role.Name)
	if err != nil {
		return nil, err
	}

	resp := map[string]string{
		"token": jwt,
	}

	return resp, nil

}

func (s *AuthService) GetUserByUuid(uuid string) (*response.User, error) {
	data, err := s.Repo.GetUserByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = response.User{
		Name:     data.Name,
		RoleName: data.Role.Name,
	}

	return &resp, nil
}

func (s *AuthService) ActivationUser(uuid string, isActivate bool) error {

	if isActivate {
		model := models.User{
			Active: true,
		}
		if err := s.Repo.UpdateUser(uuid, &model); err != nil {
			log.Println(err.Error())
			return response.SERVICE_INTERR
		}
		return nil
	}

	if err := s.Repo.UnactivateUser(uuid); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}
	return nil

}

func (s *AuthService) UpdatePassword(userUuid string, req *request.UpdatePassword) error {
	user, err := s.Repo.GetUserByUuid(userUuid)
	if err != nil {
		log.Println(err)
		return response.SERVICE_INTERR
	}

	if err := utils.ComparePassword(user.Password, req.OldPassword); err != nil {
		return response.BADREQ_ERR("Password lama salah")
	}

	if req.NewPassword != req.ConfirmPassword {
		return response.BADREQ_ERR("Konfirmasi password tidak sesuai")
	}

	model := models.User{
		Password: req.NewPassword,
	}

	if err := s.Repo.UpdateUser(userUuid, &model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil

}
