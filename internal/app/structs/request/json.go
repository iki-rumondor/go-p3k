package request

type SignIn struct {
	Username string `json:"username" valid:"required~field username tidak ditemukan"`
	Password string `json:"password" valid:"required~field password tidak ditemukan"`
}

type RegisterGuest struct {
	RoleID          uint   `json:"role_id" valid:"required~field role pengguna tidak ditemukan"`
	Fullname        string `json:"fullname" valid:"required~field nama lengkap tidak ditemukan"`
	Address         string `json:"address" valid:"required~field alamat tidak ditemukan"`
	PhoneNumber     string `json:"phone_number" valid:"required~field nomor handphone tidak ditemukan"`
	Username        string `json:"username" valid:"required~field username tidak ditemukan"`
	Password        string `json:"password" valid:"required~field password tidak ditemukan"`
	ConfirmPassword string `json:"confirm_password" valid:"required~field konfirmasi password tidak ditemukan"`
}

type Activation struct {
	Status bool `json:"status"`
}

type Category struct {
	Name string `json:"name" valid:"required~field nama tidak ditemukan"`
}

type Shop struct {
	CategoryUuid string `json:"category_uuid" valid:"required~field kategori tidak ditemukan"`
	Name         string `json:"name" valid:"required~field nama tidak ditemukan"`
	Owner        string `json:"owner" valid:"required~field nama pemilik tidak ditemukan"`
	Address      string `json:"address" valid:"required~field alamat tidak ditemukan"`
	PhoneNumber  string `json:"phone_number" valid:"required~field nomor handphone tidak ditemukan"`
}

type BuyProduct struct {
	ProductUuid string `json:"product_uuid" valid:"required~field Uuid produk tidak ditemukan"`
	Quantity    int64  `json:"quantity" valid:"required~field jumlah tidak ditemukan"`
}

type Citizen struct {
	Name        string `json:"name" valid:"required~field nama tidak ditemukan"`
	Nik         string `json:"nik" valid:"required~field nik tidak ditemukan"`
	Address     string `json:"address" valid:"required~field alamat tidak ditemukan"`
	PhoneNumber string `json:"phone_number" valid:"required~field nomor handphone tidak ditemukan"`
}

type Member struct {
	Name        string `json:"name" valid:"required~field nama tidak ditemukan"`
	IsImportant bool   `json:"is_important"`
	Position    string `json:"position" valid:"required~field jabatan tidak ditemukan"`
}

type MemberActivity struct {
	ActivityUuid string `json:"activity_uuid" valid:"required~field uuid kegiatan tidak ditemukan"`
	MemberUuid   string `json:"member_uuid" valid:"required~field uuid anggota tidak ditemukan"`
}

type UpdatePassword struct {
	OldPassword     string `json:"old_password" valid:"required~field password lama tidak ditemukan"`
	NewPassword     string `json:"new_password" valid:"required~field password baru tidak ditemukan"`
	ConfirmPassword string `json:"confirm_password" valid:"required~field konfirmasi password tidak ditemukan"`
}

type AcceptPresence struct {
	Status bool `json:"status"`
}