package services

import (
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

func (s *FetchService) GetAllProducts() (*[]response.Product, error) {
	data, err := s.Repo.GetAllProducts()
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
		Uuid:  item.Uuid,
		Name:  item.Name,
		Price: item.Price,
		Stock: item.Stock,
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
