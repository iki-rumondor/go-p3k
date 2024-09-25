package services

import (
	"strconv"

	"github.com/iki-rumondor/go-p3k/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/response"
)

type FetchService struct {
	Repo interfaces.FetchInterface
}

func NewFetchService(repo interfaces.FetchInterface) *FetchService {
	return &FetchService{
		Repo: repo,
	}
}

func (s *FetchService) GetCitizens() (*[]response.Citizen, error) {
	data, err := s.Repo.GetCitizens()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Citizen
	for _, item := range *data {
		resp = append(resp, response.Citizen{
			Uuid:        item.Uuid,
			Name:        item.Name,
			Address:     item.Address,
			PhoneNumber: item.PhoneNumber,
			Nik:         item.Nik,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			User: &response.User{
				Uuid:     item.User.Uuid,
				Username: item.User.Username,
			},
		})
	}

	return &resp, nil
}

func (s *FetchService) GetCitizenByUuid(uuid string) (*response.Citizen, error) {
	item, err := s.Repo.GetCitizenByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = response.Citizen{
		Uuid:        item.Uuid,
		Name:        item.Name,
		Address:     item.Address,
		PhoneNumber: item.PhoneNumber,
		Nik:         item.Nik,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		User: &response.User{
			Uuid:     item.User.Uuid,
			Username: item.User.Username,
		},
	}

	return &resp, nil
}

func (s *FetchService) GetGuests() (*[]response.Guest, error) {
	data, err := s.Repo.GetGuests()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Guest
	for _, item := range *data {
		resp = append(resp, response.Guest{
			Uuid:        item.Uuid,
			Name:        item.Name,
			Address:     item.Address,
			PhoneNumber: item.PhoneNumber,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			User: &response.User{
				Uuid:     item.User.Uuid,
				IsActive: item.User.Active,
				Username: item.User.Username,
			},
		})
	}

	return &resp, nil
}

func (s *FetchService) GetGuestByUuid(uuid string) (*response.Guest, error) {
	item, err := s.Repo.GetGuestByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = response.Guest{
		Uuid:        item.Uuid,
		Name:        item.Name,
		Address:     item.Address,
		PhoneNumber: item.PhoneNumber,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		User: &response.User{
			Uuid:     item.User.Uuid,
			IsActive: item.User.Active,
			Username: item.User.Username,
		},
	}

	return &resp, nil
}

func (s *FetchService) GetCategories() (*[]response.Category, error) {
	data, err := s.Repo.GetCategories()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Category
	for _, item := range *data {
		resp = append(resp, response.Category{
			Uuid:      item.Uuid,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return &resp, nil
}

func (s *FetchService) GetCategoryByUuid(uuid string) (*response.Category, error) {
	item, err := s.Repo.GetCategoryByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = response.Category{
		Uuid:      item.Uuid,
		Name:      item.Name,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}

	return &resp, nil
}

func (s *FetchService) GetShops() (*[]response.Shop, error) {
	data, err := s.Repo.GetShops()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Shop
	for _, item := range *data {
		resp = append(resp, response.Shop{
			Uuid:        item.Uuid,
			Name:        item.Name,
			Owner:       item.Owner,
			Address:     item.Address,
			PhoneNumber: item.PhoneNumber,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			User: &response.User{
				Uuid:     item.User.Uuid,
				IsActive: item.User.Active,
				Username: item.User.Username,
			},
			Category: &response.Category{
				Uuid: item.Category.Uuid,
				Name: item.Category.Name,
			},
		})
	}

	return &resp, nil
}

func (s *FetchService) GetShopByUuid(uuid string) (*response.Shop, error) {
	item, err := s.Repo.GetShopByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = response.Shop{
		Uuid:        item.Uuid,
		Name:        item.Name,
		Owner:       item.Owner,
		Address:     item.Address,
		PhoneNumber: item.PhoneNumber,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		User: &response.User{
			Uuid:     item.User.Uuid,
			IsActive: item.User.Active,
			Username: item.User.Username,
		},
		Category: &response.Category{
			Uuid: item.Category.Uuid,
			Name: item.Category.Name,
		},
	}

	return &resp, nil
}

func (s *FetchService) GetAllProducts(limit string) (*[]response.Product, error) {
	var limitNumber = -1
	if limit != "" {
		result, err := strconv.Atoi(limit)
		if err == nil {
			limitNumber = result
		}
	}

	data, err := s.Repo.GetAllProducts(limitNumber)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Product
	for _, item := range *data {
		resp = append(resp, response.Product{
			Uuid:      item.Uuid,
			Name:      item.Name,
			Price:     item.Price,
			Stock:     item.Stock,
			ImageName: item.Image,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			Shop: &response.Shop{
				Name: item.Shop.Name,
				Category: &response.Category{
					Name: item.Shop.Category.Name,
				},
			},
		})
	}

	return &resp, nil
}

func (s *FetchService) GetPublicProductByUuid(uuid string) (*response.Product, error) {
	item, err := s.Repo.GetPublicProductByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = response.Product{
		Uuid:      item.Uuid,
		Name:      item.Name,
		Price:     item.Price,
		Stock:     item.Stock,
		ImageName: item.Image,
		Shop: &response.Shop{
			Name: item.Shop.Name,
			Category: &response.Category{
				Name: item.Shop.Category.Name,
			},
		},
	}

	return &resp, nil
}

func (s *FetchService) GetProducts(userUuid string) (*[]response.Product, error) {
	data, err := s.Repo.GetProducts(userUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Product
	for _, item := range *data {
		resp = append(resp, response.Product{
			Uuid:      item.Uuid,
			Name:      item.Name,
			Price:     item.Price,
			Stock:     item.Stock,
			ImageName: item.Image,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return &resp, nil
}

func (s *FetchService) GetProductByUuid(userUuid, uuid string) (*response.Product, error) {
	item, err := s.Repo.GetProductByUuid(userUuid, uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = response.Product{
		Uuid:      item.Uuid,
		Name:      item.Name,
		Price:     item.Price,
		Stock:     item.Stock,
		ImageName: item.Image,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}

	return &resp, nil
}

func (s *FetchService) GetProductTransactions(userUuid string) (*[]response.ProductTransaction, error) {
	data, err := s.Repo.GetProductTransactions(userUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.ProductTransaction
	for _, item := range *data {
		resp = append(resp, response.ProductTransaction{
			Uuid:       item.Uuid,
			Quantity:   item.Quantity,
			IsResponse: item.IsResponse,
			IsAccept:   item.IsAccept,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			Product: &response.Product{
				Name: item.Product.Name,
			},
		})
	}

	return &resp, nil
}

func (s *FetchService) GetProductTransactionByUuid(userUuid, uuid string) (*response.ProductTransaction, error) {
	item, err := s.Repo.GetProductTransactionByUuid(userUuid, uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var phoneNumber string

	if item.User.Role.Name == "GUEST" {
		phoneNumber = item.User.Guest.PhoneNumber
	}

	if item.User.Role.Name == "CITIZEN" {
		phoneNumber = "123"
	}

	var resp = response.ProductTransaction{
		Uuid:       item.Uuid,
		Quantity:   item.Quantity,
		IsResponse: item.IsResponse,
		IsAccept:   item.IsAccept,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
		Product: &response.Product{
			Name:      item.Product.Name,
			Stock:     item.Product.Stock,
			Price:     item.Product.Price,
			ImageName: item.Product.Image,
			CreatedAt: item.Product.CreatedAt,
		},
		User: &response.User{
			Name:        item.User.Name,
			RoleName:    item.User.Role.Name,
			PhoneNumber: phoneNumber,
		},
	}

	return &resp, nil
}

func (s *FetchService) GetProductTransactionsByShop(userUuid string) (*[]response.ProductTransaction, error) {
	data, err := s.Repo.GetProductTransactionsByShop(userUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.ProductTransaction
	for _, item := range *data {
		resp = append(resp, response.ProductTransaction{
			Uuid:       item.Uuid,
			Quantity:   item.Quantity,
			IsResponse: item.IsResponse,
			IsAccept:   item.IsAccept,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			Product: &response.Product{
				Name: item.Product.Name,
			},
			User: &response.User{
				Name: item.User.Name,
			},
		})
	}

	return &resp, nil
}

func (s *FetchService) GetMembers() (*[]response.Member, error) {
	data, err := s.Repo.GetMembers()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Member
	for _, item := range *data {
		resp = append(resp, response.Member{
			Uuid:      item.Uuid,
			Name:      item.Name,
			Group:     item.Group,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			User: &response.User{
				Uuid:     item.User.Uuid,
				Username: item.User.Username,
			},
		})
	}

	return &resp, nil
}

func (s *FetchService) GetMemberByUuid(uuid string) (*response.Member, error) {
	item, err := s.Repo.GetMemberByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = response.Member{
		Uuid:      item.Uuid,
		Name:      item.Name,
		Group:     item.Group,
		Position:  item.Position,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
		User: &response.User{
			Uuid:     item.User.Uuid,
			Username: item.User.Username,
		},
	}

	return &resp, nil
}

func (s *FetchService) GetActivities(limit string) (*[]response.Activity, error) {
	var limitNumber = -1
	if limit != "" {
		result, err := strconv.Atoi(limit)
		if err == nil {
			limitNumber = result
		}
	}

	data, err := s.Repo.GetActivities(limitNumber)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Activity
	for _, item := range *data {
		resp = append(resp, response.Activity{
			Uuid:        item.Uuid,
			Title:       item.Title,
			Description: item.Description,
			ImageName:   item.ImageName,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			CreatedUser: &response.User{
				Name: item.CreatedUser.Name,
			},
			UpdatedUser: &response.User{
				Name: item.UpdatedUser.Name,
			},
		})
	}

	return &resp, nil
}

func (s *FetchService) GetActivityByUuid(uuid string) (*response.Activity, error) {
	item, err := s.Repo.GetActivityByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var members = []response.Member{}
	for _, i := range *item.Members {
		members = append(members, response.Member{
			Uuid:  i.Member.Uuid,
			Name:  i.Member.Name,
			Group: i.Member.Group,
		})
	}
	var resp = response.Activity{
		Uuid:        item.Uuid,
		Title:       item.Title,
		Group:       item.Group,
		Description: item.Description,
		ImageName:   item.ImageName,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		CreatedUser: &response.User{
			Name: item.CreatedUser.Name,
		},
		UpdatedUser: &response.User{
			Name: item.UpdatedUser.Name,
		},
		Members: &members,
	}

	return &resp, nil
}

func (s *FetchService) GetMembersNotInActivity(activityUuid string) (*[]response.Member, error) {
	data, err := s.Repo.GetMembersNotInActivity(activityUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Member
	for _, item := range *data {
		resp = append(resp, response.Member{
			Uuid: item.Uuid,
			Name: item.Name,
		})
	}

	return &resp, nil
}
