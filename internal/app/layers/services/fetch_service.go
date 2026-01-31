package services

import (
	"strconv"

	"github.com/iki-rumondor/go-p3k/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/request"
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
			Region: &response.Region{
				Uuid: item.Region.Uuid,
				Name: item.Region.Name,
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
		Region: &response.Region{
			Uuid: item.Region.Uuid,
			Name: item.Region.Name,
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

func (s *FetchService) GetShops(limit string) (*[]response.Shop, error) {
	var limitNumber int
	if limit != "" {
		result, err := strconv.Atoi(limit)
		if err == nil {
			limitNumber = result
		}
	}

	data, err := s.Repo.GetShops(limitNumber)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Shop
	for _, item := range *data {
		resp = append(resp, response.Shop{
			Uuid:          item.Uuid,
			Name:          item.Name,
			Owner:         item.Owner,
			Address:       item.Address,
			PhoneNumber:   item.PhoneNumber,
			ShopImage:     item.ShopImage,
			IdentityImage: item.IdentityImage,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
			User: &response.User{
				Uuid:     item.User.Uuid,
				IsActive: item.User.Active,
				Username: item.User.Username,
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
	}

	return &resp, nil
}

func (s *FetchService) GetAllProducts(limit, shopUuid string) (*[]response.Product, error) {
	var limitNumber = -1
	if limit != "" {
		result, err := strconv.Atoi(limit)
		if err == nil {
			limitNumber = result
		}
	}

	data, err := s.Repo.GetAllProducts(limitNumber, shopUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Product
	for _, item := range *data {
		resp = append(resp, response.Product{
			Uuid:         item.Uuid,
			Name:         item.Name,
			Price:        item.Price,
			Stock:        item.Stock,
			Unit:         item.Unit,
			ImageName:    item.Image,
			CategoryName: item.Category.Name,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
			Shop: &response.Shop{
				Name: item.Shop.Name,
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
		Uuid:         item.Uuid,
		Name:         item.Name,
		Price:        item.Price,
		Stock:        item.Stock,
		Unit:         item.Unit,
		ImageName:    item.Image,
		CategoryName: item.Category.Name,
		Shop: &response.Shop{
			Name:        item.Shop.Name,
			PhoneNumber: item.Shop.PhoneNumber,
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
			Uuid:         item.Uuid,
			CategoryName: item.Category.Name,
			Name:         item.Name,
			Price:        item.Price,
			Stock:        item.Stock,
			Unit:         item.Unit,
			ImageName:    item.Image,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
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
		Uuid:         item.Uuid,
		CategoryUuid: item.Category.Uuid,
		Name:         item.Name,
		Price:        item.Price,
		Stock:        item.Stock,
		Unit:         item.Unit,
		ImageName:    item.Image,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
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
			ProofFile:  item.ProofFile,
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
		ProofFile:  item.ProofFile,
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
		Product: &response.Product{
			Name:      item.Product.Name,
			Stock:     item.Product.Stock,
			Price:     item.Product.Price,
			Unit:      item.Product.Unit,
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

func (s *FetchService) GetProductTransactionsByShop(userUuid string, isAccept bool) (*[]response.ProductTransaction, error) {
	data, err := s.Repo.GetProductTransactionsByShop(userUuid, isAccept)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.ProductTransaction
	for _, item := range *data {
		resp = append(resp, response.ProductTransaction{
			Uuid:       item.Uuid,
			Quantity:   item.Quantity,
			Revenue:    item.Revenue,
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

func (s *FetchService) GetMembers(group string) (*[]response.Member, error) {
	data, err := s.Repo.GetMembers(group)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Member
	for _, item := range *data {
		resp = append(resp, response.Member{
			Uuid:        item.Uuid,
			Name:        item.Name,
			Position:    item.Position,
			IsImportant: item.IsImportant,
			Group:       item.Group,
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

func (s *FetchService) GetMemberByUuid(uuid string) (*response.Member, error) {
	item, err := s.Repo.GetMemberByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = response.Member{
		Uuid:        item.Uuid,
		Name:        item.Name,
		IsImportant: item.IsImportant,
		Position:    item.Position,
		Group:       item.Group,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		User: &response.User{
			Uuid:     item.User.Uuid,
			Username: item.User.Username,
		},
	}

	return &resp, nil
}

func (s *FetchService) GetActivities(limit, group, startDate, endDate string) (*[]response.Activity, error) {
	var limitNumber = -1
	if limit != "" {
		result, err := strconv.Atoi(limit)
		if err == nil {
			limitNumber = result
		}
	}

	data, err := s.Repo.GetActivitiesWithDate(limitNumber, group, startDate, endDate)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.Activity
	for _, item := range *data {
		var members []response.Member
		for _, member := range *item.Members {
			if member.IsAccept {
				members = append(members, response.Member{
					Name:     member.Member.Name,
					Position: member.Member.Position,
				})
			}
		}
		resp = append(resp, response.Activity{
			Uuid:        item.Uuid,
			Title:       item.Title,
			Group:       item.Group,
			Description: item.Description,
			ImageName:   item.ImageName,
			Location:    item.Location,
			StartTime:   item.StartTime,
			EndTime:     item.EndTime,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			CreatedUser: &response.User{
				Name: item.CreatedUser.Name,
			},
			UpdatedUser: &response.User{
				Name: item.UpdatedUser.Name,
			},
			Members: &members,
		})
	}

	return &resp, nil
}

func (s *FetchService) GetActivityByUuid(uuid string, queries request.ActivityQuery) (*response.Activity, error) {
	item, err := s.Repo.GetActivityByUuid(uuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var memberActivityResp response.MemberActivity
	if queries.Member != "" {
		memberActivity, err := s.Repo.GetMemberActivity(queries.Member, uuid)
		if err == nil {
			memberActivityResp = response.MemberActivity{
				AttendanceImage: memberActivity.AttendenceImage,
				IsAccept:        memberActivity.IsAccept,
			}
		}
	}

	var members = []response.Member{}
	for _, i := range *item.Members {
		if i.IsAccept {
			members = append(members, response.Member{
				Uuid:            i.Member.Uuid,
				AttendanceImage: i.AttendenceImage,
				Name:            i.Member.Name,
				IsImportant:     i.Member.IsImportant,
			})
		}
	}

	var resp = response.Activity{
		Uuid:        item.Uuid,
		Title:       item.Title,
		Group:       item.Group,
		Description: item.Description,
		ImageName:   item.ImageName,
		Location:    item.Location,
		StartTime:   item.StartTime,
		EndTime:     item.EndTime,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		CreatedUser: &response.User{
			Name: item.CreatedUser.Name,
		},
		UpdatedUser: &response.User{
			Name: item.UpdatedUser.Name,
		},
		Members:        &members,
		MemberActivity: &memberActivityResp,
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

func (s *FetchService) GetMemberActivity(userUuid, activityUuid string) (*response.MemberActivity, error) {
	data, err := s.Repo.GetMemberActivity(userUuid, activityUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp = response.MemberActivity{
		AttendanceImage: data.AttendenceImage,
		IsAccept:        data.IsAccept,
		Activity: &response.Activity{
			Title:       data.Activity.Title,
			Description: data.Activity.Description,
			Group:       data.Activity.Group,
			UpdatedAt:   data.Activity.UpdatedAt,
			UpdatedUser: &response.User{
				Name: data.Activity.UpdatedUser.Name,
			},
		},
	}

	return &resp, nil
}

func (s *FetchService) GetAdminDashboard() (*response.AdminDashboard, error) {
	guestsInactive, err := s.Repo.CountGuestsInactive()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	shopsInactive, err := s.Repo.CountShopsInactive()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	resp := response.AdminDashboard{
		GuestsInactive: guestsInactive,
		ShopsInactive:  shopsInactive,
	}

	return &resp, nil
}

func (s *FetchService) GetShopDashboard(userUuid string) (*response.ShopDashboard, error) {
	products, err := s.Repo.CountShopProducts(userUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	transactions, err := s.Repo.CountShopUnprocessTransactions(userUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	resp := response.ShopDashboard{
		Products:              products,
		UnprocessTransactions: transactions,
	}

	return &resp, nil
}

func (s *FetchService) GetShopByUser(userUuid string) (*response.Shop, error) {
	shop, err := s.Repo.GetShopByUser(userUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	resp := response.Shop{
		Uuid:        shop.Uuid,
		Name:        shop.Name,
		Owner:       shop.Owner,
		Address:     shop.Address,
		PhoneNumber: shop.PhoneNumber,
	}

	return &resp, nil
}

func (s *FetchService) GetGuestDashboard(userUuid string) (*response.GuestDashboard, error) {
	guest, err := s.Repo.GetGuestByUser(userUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	unprocess_transactions, err := s.Repo.CountUserUnprocessTransactions(userUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	success_transactions, err := s.Repo.CountUserSuccessTransactions(userUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	resp := response.GuestDashboard{
		UnprocessTransactions: unprocess_transactions,
		SuccessTransactions:   success_transactions,
		Uuid:                  guest.Uuid,
		Name:                  guest.Name,
		Address:               guest.Address,
		PhoneNumber:           guest.PhoneNumber,
	}

	return &resp, nil
}

func (s *FetchService) GetMemberDashboard(userUuid string) (*response.MemberDashboard, error) {
	member, err := s.Repo.GetMemberByUserUuid(userUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	activities, err := s.Repo.CountActivities()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	resp := response.MemberDashboard{
		Activities: activities,
		Position:   member.Position,
	}

	return &resp, nil
}

func (s *FetchService) GetMemberActivities(userUuid string) (*response.MemberActivities, error) {
	member, err := s.Repo.GetMemberByUserUuid(userUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	activities, err := s.Repo.GetActivities(-1, "")
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var activitiesResp []response.Activity
	for _, item := range *activities {
		activitiesResp = append(activitiesResp, response.Activity{
			Uuid:      item.Uuid,
			Title:     item.Title,
			ImageName: item.ImageName,
			CreatedAt: item.CreatedAt,
			CreatedUser: &response.User{
				Name: item.CreatedUser.Name,
			},
		})
	}

	resp := response.MemberActivities{
		Activities:  &activitiesResp,
		IsImportant: member.IsImportant,
		IsHeadgroup: member.IsHeadgroup,
	}

	return &resp, nil
}

func (s *FetchService) GetMemberByUser(userUuid string) (*response.Member, error) {
	member, err := s.Repo.GetMemberByUserUuid(userUuid)
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	resp := response.Member{
		IsImportant: member.IsImportant,
		IsHeadgroup: member.IsHeadgroup,
	}

	return &resp, nil
}

func (s *FetchService) GetAllMemberActivities() (*[]response.MemberActivity, error) {
	results, err := s.Repo.GetMemberActivities()
	if err != nil {
		return nil, response.SERVICE_INTERR
	}

	var resp []response.MemberActivity

	for _, item := range *results {
		resp = append(resp, response.MemberActivity{
			Uuid:            item.Uuid,
			AttendanceImage: item.AttendenceImage,
			IsAccept:        item.IsAccept,
			CreatedAt:       item.CreatedAt,
			Member: &response.Member{
				Name: item.Member.Name,
			},
			Activity: &response.Activity{
				Title: item.Activity.Title,
			},
		})
	}

	return &resp, nil
}
