package response

type User struct {
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Username string `json:"username"`
	RoleName string `json:"role_name"`
	IsActive bool   `json:"is_active"`
}

type Guest struct {
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	User        *User  `json:"user"`
}

type Category struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type Shop struct {
	Uuid        string    `json:"uuid"`
	Name        string    `json:"name"`
	Owner       string    `json:"owner"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   int64     `json:"created_at"`
	UpdatedAt   int64     `json:"updated_at"`
	User        *User     `json:"user"`
	Category    *Category `json:"category"`
}

type Product struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Stock     int64  `json:"stock"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Shop      *Shop  `json:"shop"`
}
