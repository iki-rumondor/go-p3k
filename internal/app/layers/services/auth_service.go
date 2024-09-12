package services

import (
	"errors"
	"log"

	"github.com/iki-rumondor/go-p3k/internal/app/layers/interfaces"
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
