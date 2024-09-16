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

func (s *AuthService) CreateGuest(req *request.RegisterGuest) error {
	if req.Password != req.ConfirmPassword {
		return response.BADREQ_ERR("Konfirmasi password tidak sama dengan password")
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
			RoleID:   4,
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
