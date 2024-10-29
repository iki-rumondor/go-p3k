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
	Uuid         string `json:"uuid"`
	CategoryUuid string `json:"category_uuid"`
	CategoryName string `json:"category_name"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Stock        int64  `json:"stock"`
	Unit         string `json:"unit"`
	ImageName    string `json:"image_name"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	Shop         *Shop  `json:"shop"`
}

type ProductTransaction struct {
	Uuid       string   `json:"uuid"`
	Quantity   int64    `json:"quantity"`
	IsResponse bool     `json:"is_response"`
	IsAccept   bool     `json:"is_accept"`
	Revenue    int64    `json:"revenue"`
	ProofFile  string   `json:"proof_file"`
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
	Uuid            string `json:"uuid"`
	Name            string `json:"name"`
	AttendanceImage string `json:"attendance_image"`
	IsImportant     bool   `json:"is_important"`
	IsHeadgroup     bool   `json:"is_headgroup"`
	Position        string `json:"position"`
	Group           int64  `json:"group"`
	CreatedAt       int64  `json:"created_at"`
	UpdatedAt       int64  `json:"updated_at"`
	User            *User  `json:"user"`
}

type Activity struct {
	Uuid           string          `json:"uuid"`
	Title          string          `json:"title"`
	Group          uint            `json:"group"`
	Description    string          `json:"description"`
	ImageName      string          `json:"image_name"`
	Location       string          `json:"location"`
	Date           int64           `json:"date"`
	CreatedAt      int64           `json:"created_at"`
	UpdatedAt      int64           `json:"updated_at"`
	CreatedUser    *User           `json:"created_user"`
	UpdatedUser    *User           `json:"updated_user"`
	Members        *[]Member       `json:"members"`
	MemberActivity *MemberActivity `json:"member_activity"`
}

type MemberActivity struct {
	Uuid            string    `json:"uuid"`
	AttendanceImage string    `json:"attendance_image"`
	IsAccept        bool      `json:"is_accept"`
	CreatedAt       int64     `json:"created_at"`
	UpdatedAt       int64     `json:"updated_at"`
	Activity        *Activity `json:"activity"`
	Member          *Member   `json:"member"`
}

type AdminDashboard struct {
	GuestsInactive int64 `json:"guests_inactive"`
	ShopsInactive  int64 `json:"shops_inactive"`
}

type ShopDashboard struct {
	Products              int64 `json:"products"`
	UnprocessTransactions int64 `json:"unprocess_transactions"`
}

type GuestDashboard struct {
	SuccessTransactions   int64 `json:"success_transactions"`
	UnprocessTransactions int64 `json:"unprocess_transactions"`
}

type MemberDashboard struct {
	Activities int64  `json:"activities"`
	Position   string `json:"position"`
}

type MemberActivities struct {
	IsHeadgroup bool        `json:"is_headgroup"`
	IsImportant bool        `json:"is_important"`
	Activities  *[]Activity `json:"activities"`
}
