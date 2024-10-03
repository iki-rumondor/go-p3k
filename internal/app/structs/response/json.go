package response

type User struct {
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	RoleName    string `json:"role_name"`
	PhoneNumber string `json:"phone_number"`
	IsActive    bool   `json:"is_active"`
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
	Uuid          string `json:"uuid"`
	Name          string `json:"name"`
	Owner         string `json:"owner"`
	Address       string `json:"address"`
	PhoneNumber   string `json:"phone_number"`
	ShopImage     string `json:"shop_image"`
	IdentityImage string `json:"identity_image"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
	User          *User  `json:"user"`
}

type Product struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Stock     int64  `json:"stock"`
	ImageName string `json:"image_name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Shop      *Shop  `json:"shop"`
}

type ProductTransaction struct {
	Uuid       string   `json:"uuid"`
	Quantity   int64    `json:"quantity"`
	IsResponse bool     `json:"is_response"`
	IsAccept   bool     `json:"is_accept"`
	CreatedAt  int64    `json:"created_at"`
	UpdatedAt  int64    `json:"updated_at"`
	User       *User    `json:"user"`
	Product    *Product `json:"product"`
}

type Citizen struct {
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Nik         string `json:"nik"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	User        *User  `json:"user"`
}

type Member struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Group     uint   `json:"group"`
	Position  string `json:"position"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	User      *User  `json:"user"`
}

type Activity struct {
	Uuid        string    `json:"uuid"`
	Title       string    `json:"title"`
	Group       uint      `json:"group"`
	Description string    `json:"description"`
	ImageName   string    `json:"image_name"`
	CreatedAt   int64     `json:"created_at"`
	UpdatedAt   int64     `json:"updated_at"`
	CreatedUser *User     `json:"created_user"`
	UpdatedUser *User     `json:"updated_user"`
	Members     *[]Member `json:"members"`
}
